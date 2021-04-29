package api

import (
	"github.com/labstack/echo/v4"
	"github.com/tidepool-org/clinic/clinics"
	"github.com/tidepool-org/clinic/patients"
	"github.com/tidepool-org/clinic/store"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

func (h *Handler) ListPatients(ec echo.Context, clinicId ClinicId, params ListPatientsParams) error {
	ctx := ec.Request().Context()
	page := pagination(params.Offset, params.Limit)
	filter := patients.Filter{
		ClinicId: strp(string(clinicId)),
		Search:   searchToString(params.Search),
	}

	list, err := h.patients.List(ctx, &filter, page)
	if err != nil {
		return err
	}

	return ec.JSON(http.StatusOK, NewPatientsDto(list))
}

func (h *Handler) CreatePatientAccount(ec echo.Context, clinicId ClinicId) error {
	ctx := ec.Request().Context()
	dto := Patient{}
	if err := ec.Bind(&dto); err != nil {
		return err
	}

	clinicObjId, err := primitive.ObjectIDFromHex(string(clinicId))
	if err != nil {
		return err
	}

	patient := NewPatient(dto)
	patient.ClinicId = &clinicObjId
	patient.Permissions = &patients.CustodialAccountPermissions

	result, err := h.patients.Create(ctx, patient)
	if err != nil {
		return err
	}

	return ec.JSON(http.StatusOK, NewPatientDto(result))
}

func (h *Handler) GetPatient(ec echo.Context, clinicId ClinicId, patientId PatientId) error {
	ctx := ec.Request().Context()
	patient, err := h.patients.Get(ctx, string(clinicId), string(patientId))
	if err != nil {
		return err
	}

	return ec.JSON(http.StatusOK, NewPatientDto(patient))
}

func (h *Handler) CreatePatientFromUser(ec echo.Context, clinicId ClinicId, patientId PatientId) error {
	ctx := ec.Request().Context()
	dto := Patient{}
	if err := ec.Bind(&dto); err != nil {
		return err
	}

	clinicObjId, err := primitive.ObjectIDFromHex(string(clinicId))
	if err != nil {
		return err
	}

	patient := patients.Patient{
		UserId: strp(string(patientId)),
		ClinicId: &clinicObjId,
		Permissions: NewPermissions(dto.Permissions),
	}
	if err = h.users.GetPatientFromExistingUser(ctx, &patient); err != nil {
		return err
	}

	result, err := h.patients.Create(ctx, patient)
	if err != nil {
		return err
	}

	return ec.JSON(http.StatusOK, NewPatientDto(result))
}

func (h *Handler) UpdatePatient(ec echo.Context, clinicId ClinicId, patientId PatientId) error {
	ctx := ec.Request().Context()
	dto := Patient{}
	if err := ec.Bind(&dto); err != nil {
		return err
	}

	patient, err := h.patients.Update(ctx, string(clinicId), string(patientId), NewPatient(dto))
	if err != nil {
		return err
	}

	return ec.JSON(http.StatusOK, NewPatientDto(patient))
}

func (h *Handler) UpdatePatientPermissions(ec echo.Context, clinicId ClinicId, patientId PatientId) error {
	ctx := ec.Request().Context()
	dto := PatientPermissions{}
	if err := ec.Bind(&dto); err != nil {
		return err
	}

	patient, err := h.patients.UpdatePermissions(ctx, string(clinicId), string(patientId), NewPermissions(&dto))
	if err != nil {
		return err
	}

	return ec.JSON(http.StatusOK, NewPatientDto(patient).Permissions)
}

func (h *Handler) ListClinicsForPatient(ec echo.Context, patientId UserId, params ListClinicsForPatientParams) error {
	ctx := ec.Request().Context()
	page := pagination(params.Offset, params.Limit)
	patientList, err := h.patients.List(ctx, &patients.Filter{UserId: strp(string(patientId))}, page)
	if err != nil {
		return err
	}

	var clinicIds []string
	for _, patient := range patientList {
		clinicIds = append(clinicIds, patient.ClinicId.Hex())
	}

	clinicList, err := h.clinics.List(ctx, &clinics.Filter{Ids: clinicIds}, store.Pagination{})
	dtos, err := NewPatientClinicRelationshipsDto(patientList, clinicList)
	if err != nil {
		return err
	}

	return ec.JSON(http.StatusOK, dtos)
}
