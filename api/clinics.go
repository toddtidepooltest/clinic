package api

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/tidepool-org/clinic/store"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
)

type ClinicServer struct {
	Store store.StorageInterface
}

type ClinicExtraFields struct {
	Active bool   `json:"active" bson:"active"`
}

type FullClinic struct {
	Clinic `bson:",inline"`
	ClinicExtraFields `bson:",inline"`
}

type FullNewClinic struct {
	NewClinic `bson:",inline"`
	ClinicExtraFields `bson:",inline"`
}
// getClinic
// (GET /clinics)
func (c *ClinicServer) GetClinics(ctx echo.Context, params GetClinicsParams) error {
	newClinic := NewClinic{}
	filter := FullNewClinic{ClinicExtraFields: ClinicExtraFields{Active: true}, NewClinic: newClinic}
	pagingParams := store.DefaultPagingParams
	if params.Limit != nil {
		pagingParams.Limit = int64(*params.Limit)
	}
	if params.Offset != nil {
		pagingParams.Offset = int64(*params.Offset)
	}

	var clinics []Clinic
	if err := c.Store.Find(store.ClinicsCollection, filter, &pagingParams, &clinics); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "error finding clinic")
	}

	return ctx.JSON(http.StatusOK, &clinics)
}

// createClinic
// (POST /clinics)
func (c *ClinicServer) PostClinics(ctx echo.Context) error {
	var newClinic FullNewClinic

	log.Printf("Entering post clincs")
	if err := ctx.Bind(&newClinic); err != nil {
		log.Printf("Format failed for post clinic body")
		return echo.NewHTTPError(http.StatusBadRequest, "error parsing parameters")
	}
	log.Printf("Middle of post clinics")
	newClinic.Active = true

	var clinicsClinicians FullClinicsClinicians
	if newID, err := c.Store.InsertOne(store.ClinicsCollection, newClinic); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error inserting clinic")
	} else {
		clinicsClinicians.ClinicId = newID
	}

	// We also have to add an initial admin - the user
	userId := ctx.Request().Header.Get(TidepoolUserIdHeaderKey)
	clinicsClinicians.Active = true
	clinicsClinicians.ClinicianId = userId
	clinicsClinicians.Permissions = &[]string{"CLINIC_ADMIN"}
	log.Printf("End of post clinics")
	if newID, err := c.Store.InsertOne(store.ClinicsCliniciansCollection, clinicsClinicians); err != nil {
		log.Printf("Wrong way of post clinics")
		return echo.NewHTTPError(http.StatusInternalServerError, "error inserting to clinician")
	} else {
		var ret struct {
			id  string    `json:"id"`
		}
		ret.id = *newID
		log.Printf("Returning from /clinics")
		return ctx.JSON(http.StatusOK, &ret)
	}
}

// (DELETE /clinic/{clinicid})
func (c *ClinicServer) DeleteClinicsClinicid(ctx echo.Context, clinicid string) error {
	objID, _ := primitive.ObjectIDFromHex(clinicid)
	filter := bson.D{{"_id", objID}}
	activeObj := bson.D{
		{"$set", bson.D{{"active", false}}},
	}
	if err := c.Store.UpdateOne(store.ClinicsCollection, filter, activeObj); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "error deleting clinic from database")
	}
	return ctx.JSON(http.StatusOK, nil)
}

// getClinic
// (GET /clinic/{clinicid})
func (c *ClinicServer) GetClinicsClinicid(ctx echo.Context, clinicid string) error {
	var clinic Clinic
	log.Printf("Get Clinic by id - id: %s", clinicid)
	objID, _ := primitive.ObjectIDFromHex(clinicid)
	filter := bson.M{"_id": objID, "active": true}
	if err := c.Store.FindOne(store.ClinicsCollection, filter, &clinic); err != nil {
		fmt.Println("Find One error ", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "error finding clinic")
	}

	return ctx.JSON(http.StatusOK, &clinic)
}

// (PATCH /clinic/{clinicid})
func (c *ClinicServer) PatchClinicsClinicid(ctx echo.Context, clinicid string) error {
	var newClinic NewClinic

	if err := ctx.Bind(&newClinic); err != nil {
		log.Printf("Format failed for patch clinic body")
		return echo.NewHTTPError(http.StatusBadRequest, "error parsing parameters")
	}
	objID, _ := primitive.ObjectIDFromHex(clinicid)
	filter := bson.D{{"_id", objID}}
	patchObj := bson.D{
		{"$set", newClinic },
	}
	if err := c.Store.UpdateOne(store.ClinicsCollection, filter, patchObj); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "error updating clinic")
	}
	return ctx.JSON(http.StatusOK, nil)
}

