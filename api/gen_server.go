// Package Clinic provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package api

import (
	"fmt"
	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/labstack/echo/v4"
	"net/http"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// getClinic
	// (GET /clinic)
	GetClinic(ctx echo.Context) error
	// createClinic
	// (POST /clinic)
	PostClinic(ctx echo.Context) error

	// (DELETE /clinic/{clinicid})
	DeleteClinicClinicid(ctx echo.Context, clinicid string) error
	// getClinic
	// (GET /clinic/{clinicid})
	GetClinicClinicid(ctx echo.Context, clinicid string) error

	// (PATCH /clinic/{clinicid})
	PatchClinicClinicid(ctx echo.Context, clinicid string) error
	// DeleteClinicianForClinic
	// (DELETE /clinic/{clinicid}/clinician/{clinicianid})
	DeleteClinicClinicidClinicianClinicianid(ctx echo.Context, clinicid string, clinicianid string) error

	// (PATCH /clinic/{clinicid}/clinician/{clinicianid})
	PatchClinicClinicidClinicianClinicianid(ctx echo.Context, clinicid string, clinicianid string) error
	// deletePatientFromClinic
	// (DELETE /clinic/{clinicid}/patient/{patientid})
	DeleteClinicClinicidPatientPatientid(ctx echo.Context, clinicid string, patientid string) error
	// addPatientToClinic
	// (PATCH /clinic/{clinicid}/patient/{patientid})
	PatchClinicClinicidPatientPatientid(ctx echo.Context, clinicid string, patientid string) error

	// (POST /clinic/{clinicid}/patient/{patientid})
	PostClinicClinicidPatientPatientid(ctx echo.Context, clinicid string, patientid string) error
	// getCliniciansForClinic
	// (GET /clinic/{clinicid}/patients)
	GetClinicClinicidPatients(ctx echo.Context, clinicid string) error
	// getClinic
	// (GET /clinics)
	GetClinics(ctx echo.Context) error
	// getCliniciansForClinic
	// (GET /clinics/{clinicid}/clinicians)
	GetClinicsClinicidClinicians(ctx echo.Context, clinicid string) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetClinic converts echo context to params.
func (w *ServerInterfaceWrapper) GetClinic(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetClinic(ctx)
	return err
}

// PostClinic converts echo context to params.
func (w *ServerInterfaceWrapper) PostClinic(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PostClinic(ctx)
	return err
}

// DeleteClinicClinicid converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteClinicClinicid(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "clinicid" -------------
	var clinicid string

	err = runtime.BindStyledParameter("simple", false, "clinicid", ctx.Param("clinicid"), &clinicid)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter clinicid: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.DeleteClinicClinicid(ctx, clinicid)
	return err
}

// GetClinicClinicid converts echo context to params.
func (w *ServerInterfaceWrapper) GetClinicClinicid(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "clinicid" -------------
	var clinicid string

	err = runtime.BindStyledParameter("simple", false, "clinicid", ctx.Param("clinicid"), &clinicid)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter clinicid: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetClinicClinicid(ctx, clinicid)
	return err
}

// PatchClinicClinicid converts echo context to params.
func (w *ServerInterfaceWrapper) PatchClinicClinicid(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "clinicid" -------------
	var clinicid string

	err = runtime.BindStyledParameter("simple", false, "clinicid", ctx.Param("clinicid"), &clinicid)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter clinicid: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PatchClinicClinicid(ctx, clinicid)
	return err
}

// DeleteClinicClinicidClinicianClinicianid converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteClinicClinicidClinicianClinicianid(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "clinicid" -------------
	var clinicid string

	err = runtime.BindStyledParameter("simple", false, "clinicid", ctx.Param("clinicid"), &clinicid)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter clinicid: %s", err))
	}

	// ------------- Path parameter "clinicianid" -------------
	var clinicianid string

	err = runtime.BindStyledParameter("simple", false, "clinicianid", ctx.Param("clinicianid"), &clinicianid)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter clinicianid: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.DeleteClinicClinicidClinicianClinicianid(ctx, clinicid, clinicianid)
	return err
}

// PatchClinicClinicidClinicianClinicianid converts echo context to params.
func (w *ServerInterfaceWrapper) PatchClinicClinicidClinicianClinicianid(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "clinicid" -------------
	var clinicid string

	err = runtime.BindStyledParameter("simple", false, "clinicid", ctx.Param("clinicid"), &clinicid)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter clinicid: %s", err))
	}

	// ------------- Path parameter "clinicianid" -------------
	var clinicianid string

	err = runtime.BindStyledParameter("simple", false, "clinicianid", ctx.Param("clinicianid"), &clinicianid)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter clinicianid: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PatchClinicClinicidClinicianClinicianid(ctx, clinicid, clinicianid)
	return err
}

// DeleteClinicClinicidPatientPatientid converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteClinicClinicidPatientPatientid(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "clinicid" -------------
	var clinicid string

	err = runtime.BindStyledParameter("simple", false, "clinicid", ctx.Param("clinicid"), &clinicid)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter clinicid: %s", err))
	}

	// ------------- Path parameter "patientid" -------------
	var patientid string

	err = runtime.BindStyledParameter("simple", false, "patientid", ctx.Param("patientid"), &patientid)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter patientid: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.DeleteClinicClinicidPatientPatientid(ctx, clinicid, patientid)
	return err
}

// PatchClinicClinicidPatientPatientid converts echo context to params.
func (w *ServerInterfaceWrapper) PatchClinicClinicidPatientPatientid(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "clinicid" -------------
	var clinicid string

	err = runtime.BindStyledParameter("simple", false, "clinicid", ctx.Param("clinicid"), &clinicid)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter clinicid: %s", err))
	}

	// ------------- Path parameter "patientid" -------------
	var patientid string

	err = runtime.BindStyledParameter("simple", false, "patientid", ctx.Param("patientid"), &patientid)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter patientid: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PatchClinicClinicidPatientPatientid(ctx, clinicid, patientid)
	return err
}

// PostClinicClinicidPatientPatientid converts echo context to params.
func (w *ServerInterfaceWrapper) PostClinicClinicidPatientPatientid(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "clinicid" -------------
	var clinicid string

	err = runtime.BindStyledParameter("simple", false, "clinicid", ctx.Param("clinicid"), &clinicid)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter clinicid: %s", err))
	}

	// ------------- Path parameter "patientid" -------------
	var patientid string

	err = runtime.BindStyledParameter("simple", false, "patientid", ctx.Param("patientid"), &patientid)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter patientid: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PostClinicClinicidPatientPatientid(ctx, clinicid, patientid)
	return err
}

// GetClinicClinicidPatients converts echo context to params.
func (w *ServerInterfaceWrapper) GetClinicClinicidPatients(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "clinicid" -------------
	var clinicid string

	err = runtime.BindStyledParameter("simple", false, "clinicid", ctx.Param("clinicid"), &clinicid)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter clinicid: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetClinicClinicidPatients(ctx, clinicid)
	return err
}

// GetClinics converts echo context to params.
func (w *ServerInterfaceWrapper) GetClinics(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetClinics(ctx)
	return err
}

// GetClinicsClinicidClinicians converts echo context to params.
func (w *ServerInterfaceWrapper) GetClinicsClinicidClinicians(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "clinicid" -------------
	var clinicid string

	err = runtime.BindStyledParameter("simple", false, "clinicid", ctx.Param("clinicid"), &clinicid)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter clinicid: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetClinicsClinicidClinicians(ctx, clinicid)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET("/clinic", wrapper.GetClinic)
	router.POST("/clinic", wrapper.PostClinic)
	router.DELETE("/clinic/:clinicid", wrapper.DeleteClinicClinicid)
	router.GET("/clinic/:clinicid", wrapper.GetClinicClinicid)
	router.PATCH("/clinic/:clinicid", wrapper.PatchClinicClinicid)
	router.DELETE("/clinic/:clinicid/clinician/:clinicianid", wrapper.DeleteClinicClinicidClinicianClinicianid)
	router.PATCH("/clinic/:clinicid/clinician/:clinicianid", wrapper.PatchClinicClinicidClinicianClinicianid)
	router.DELETE("/clinic/:clinicid/patient/:patientid", wrapper.DeleteClinicClinicidPatientPatientid)
	router.PATCH("/clinic/:clinicid/patient/:patientid", wrapper.PatchClinicClinicidPatientPatientid)
	router.POST("/clinic/:clinicid/patient/:patientid", wrapper.PostClinicClinicidPatientPatientid)
	router.GET("/clinic/:clinicid/patients", wrapper.GetClinicClinicidPatients)
	router.GET("/clinics", wrapper.GetClinics)
	router.GET("/clinics/:clinicid/clinicians", wrapper.GetClinicsClinicidClinicians)

}

