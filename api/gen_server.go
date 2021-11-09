// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.9.0 DO NOT EDIT.
package api

import (
	"fmt"
	"net/http"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/labstack/echo/v4"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// List Clinics for Clinician
	// (GET /v1/clinicians/{userId}/clinics)
	ListClinicsForClinician(ctx echo.Context, userId UserId, params ListClinicsForClinicianParams) error
	// Enable Clinics
	// (POST /v1/clinicians/{userId}/migrate)
	EnableNewClinicExperience(ctx echo.Context, userId string) error
	// List Clinics
	// (GET /v1/clinics)
	ListClinics(ctx echo.Context, params ListClinicsParams) error
	// Create Clinic
	// (POST /v1/clinics)
	CreateClinic(ctx echo.Context) error
	// Get Clinic by Share Code
	// (GET /v1/clinics/share_code/{shareCode})
	GetClinicByShareCode(ctx echo.Context, shareCode string) error
	// Get Clinic
	// (GET /v1/clinics/{clinicId})
	GetClinic(ctx echo.Context, clinicId ClinicId) error
	// Update Clinic
	// (PUT /v1/clinics/{clinicId})
	UpdateClinic(ctx echo.Context, clinicId ClinicId) error
	// List Clinicians
	// (GET /v1/clinics/{clinicId}/clinicians)
	ListClinicians(ctx echo.Context, clinicId ClinicId, params ListCliniciansParams) error
	// Create Clinician
	// (POST /v1/clinics/{clinicId}/clinicians)
	CreateClinician(ctx echo.Context, clinicId ClinicId) error
	// Delete Clinician
	// (DELETE /v1/clinics/{clinicId}/clinicians/{clinicianId})
	DeleteClinician(ctx echo.Context, clinicId ClinicId, clinicianId ClinicianId) error
	// Get Clinician
	// (GET /v1/clinics/{clinicId}/clinicians/{clinicianId})
	GetClinician(ctx echo.Context, clinicId ClinicId, clinicianId ClinicianId) error
	// Update Clinician
	// (PUT /v1/clinics/{clinicId}/clinicians/{clinicianId})
	UpdateClinician(ctx echo.Context, clinicId ClinicId, clinicianId ClinicianId) error
	// Delete Invited Clinician
	// (DELETE /v1/clinics/{clinicId}/invites/clinicians/{inviteId}/clinician)
	DeleteInvitedClinician(ctx echo.Context, clinicId ClinicId, inviteId InviteId) error
	// Get Invited Clinician
	// (GET /v1/clinics/{clinicId}/invites/clinicians/{inviteId}/clinician)
	GetInvitedClinician(ctx echo.Context, clinicId ClinicId, inviteId InviteId) error
	// Associate Clinician to User
	// (PATCH /v1/clinics/{clinicId}/invites/clinicians/{inviteId}/clinician)
	AssociateClinicianToUser(ctx echo.Context, clinicId ClinicId, inviteId InviteId) error
	// Trigger initial migration
	// (POST /v1/clinics/{clinicId}/migrate)
	TriggerInitialMigration(ctx echo.Context, clinicId string) error
	// Retrieve Migration Status
	// (GET /v1/clinics/{clinicId}/migrations)
	ListMigrations(ctx echo.Context, clinicId string) error
	// Migrate Legacy Clinician Patients
	// (POST /v1/clinics/{clinicId}/migrations)
	MigrateLegacyClinicianPatients(ctx echo.Context, clinicId string) error
	// Get Migration
	// (GET /v1/clinics/{clinicId}/migrations/{userId})
	GetMigration(ctx echo.Context, clinicId Id, userId UserId) error
	// Update Migration
	// (PATCH /v1/clinics/{clinicId}/migrations/{userId})
	UpdateMigration(ctx echo.Context, clinicId Id, userId UserId) error
	// List Patients
	// (GET /v1/clinics/{clinicId}/patients)
	ListPatients(ctx echo.Context, clinicId ClinicId, params ListPatientsParams) error
	// Create Patient Account
	// (POST /v1/clinics/{clinicId}/patients)
	CreatePatientAccount(ctx echo.Context, clinicId ClinicId) error
	// Delete Patient
	// (DELETE /v1/clinics/{clinicId}/patients/{patientId})
	DeletePatient(ctx echo.Context, clinicId ClinicId, patientId PatientId) error
	// Get Patient
	// (GET /v1/clinics/{clinicId}/patients/{patientId})
	GetPatient(ctx echo.Context, clinicId ClinicId, patientId PatientId) error
	// Create Patient from Existing User
	// (POST /v1/clinics/{clinicId}/patients/{patientId})
	CreatePatientFromUser(ctx echo.Context, clinicId ClinicId, patientId PatientId) error
	// Update Patient
	// (PUT /v1/clinics/{clinicId}/patients/{patientId})
	UpdatePatient(ctx echo.Context, clinicId ClinicId, patientId PatientId) error
	// Update Patient Permissions
	// (PUT /v1/clinics/{clinicId}/patients/{patientId}/permissions)
	UpdatePatientPermissions(ctx echo.Context, clinicId ClinicId, patientId PatientId) error
	// Delete Patient Permission
	// (DELETE /v1/clinics/{clinicId}/patients/{patientId}/permissions/{permission})
	DeletePatientPermission(ctx echo.Context, clinicId ClinicId, patientId PatientId, permission string) error
	// List Clinics for Patient
	// (GET /v1/patients/{userId}/clinics)
	ListClinicsForPatient(ctx echo.Context, userId UserId, params ListClinicsForPatientParams) error

	// (DELETE /v1/users/{userId}/clinics)
	DeleteUserFromClinics(ctx echo.Context, userId UserId) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// ListClinicsForClinician converts echo context to params.
func (w *ServerInterfaceWrapper) ListClinicsForClinician(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "userId" -------------
	var userId UserId

	err = runtime.BindStyledParameterWithLocation("simple", false, "userId", runtime.ParamLocationPath, ctx.Param("userId"), &userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter userId: %s", err))
	}

	ctx.Set(SessionTokenScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params ListClinicsForClinicianParams
	// ------------- Optional query parameter "offset" -------------

	err = runtime.BindQueryParameter("form", true, false, "offset", ctx.QueryParams(), &params.Offset)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter offset: %s", err))
	}

	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", ctx.QueryParams(), &params.Limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter limit: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.ListClinicsForClinician(ctx, userId, params)
	return err
}

// EnableNewClinicExperience converts echo context to params.
func (w *ServerInterfaceWrapper) EnableNewClinicExperience(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "userId" -------------
	var userId string

	err = runtime.BindStyledParameterWithLocation("simple", false, "userId", runtime.ParamLocationPath, ctx.Param("userId"), &userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter userId: %s", err))
	}

	ctx.Set(SessionTokenScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.EnableNewClinicExperience(ctx, userId)
	return err
}

// ListClinics converts echo context to params.
func (w *ServerInterfaceWrapper) ListClinics(ctx echo.Context) error {
	var err error

	ctx.Set(SessionTokenScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params ListClinicsParams
	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", ctx.QueryParams(), &params.Limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter limit: %s", err))
	}

	// ------------- Optional query parameter "offset" -------------

	err = runtime.BindQueryParameter("form", true, false, "offset", ctx.QueryParams(), &params.Offset)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter offset: %s", err))
	}

	// ------------- Optional query parameter "shareCode" -------------

	err = runtime.BindQueryParameter("form", true, false, "shareCode", ctx.QueryParams(), &params.ShareCode)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter shareCode: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.ListClinics(ctx, params)
	return err
}

// CreateClinic converts echo context to params.
func (w *ServerInterfaceWrapper) CreateClinic(ctx echo.Context) error {
	var err error

	ctx.Set(SessionTokenScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.CreateClinic(ctx)
	return err
}

// GetClinicByShareCode converts echo context to params.
func (w *ServerInterfaceWrapper) GetClinicByShareCode(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "shareCode" -------------
	var shareCode string

	err = runtime.BindStyledParameterWithLocation("simple", false, "shareCode", runtime.ParamLocationPath, ctx.Param("shareCode"), &shareCode)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter shareCode: %s", err))
	}

	ctx.Set(SessionTokenScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetClinicByShareCode(ctx, shareCode)
	return err
}

// GetClinic converts echo context to params.
func (w *ServerInterfaceWrapper) GetClinic(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "clinicId" -------------
	var clinicId ClinicId

	err = runtime.BindStyledParameterWithLocation("simple", false, "clinicId", runtime.ParamLocationPath, ctx.Param("clinicId"), &clinicId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter clinicId: %s", err))
	}

	ctx.Set(SessionTokenScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetClinic(ctx, clinicId)
	return err
}

// UpdateClinic converts echo context to params.
func (w *ServerInterfaceWrapper) UpdateClinic(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "clinicId" -------------
	var clinicId ClinicId

	err = runtime.BindStyledParameterWithLocation("simple", false, "clinicId", runtime.ParamLocationPath, ctx.Param("clinicId"), &clinicId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter clinicId: %s", err))
	}

	ctx.Set(SessionTokenScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.UpdateClinic(ctx, clinicId)
	return err
}

// ListClinicians converts echo context to params.
func (w *ServerInterfaceWrapper) ListClinicians(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "clinicId" -------------
	var clinicId ClinicId

	err = runtime.BindStyledParameterWithLocation("simple", false, "clinicId", runtime.ParamLocationPath, ctx.Param("clinicId"), &clinicId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter clinicId: %s", err))
	}

	ctx.Set(SessionTokenScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params ListCliniciansParams
	// ------------- Optional query parameter "search" -------------

	err = runtime.BindQueryParameter("form", true, false, "search", ctx.QueryParams(), &params.Search)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter search: %s", err))
	}

	// ------------- Optional query parameter "offset" -------------

	err = runtime.BindQueryParameter("form", true, false, "offset", ctx.QueryParams(), &params.Offset)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter offset: %s", err))
	}

	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", ctx.QueryParams(), &params.Limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter limit: %s", err))
	}

	// ------------- Optional query parameter "email" -------------

	err = runtime.BindQueryParameter("form", true, false, "email", ctx.QueryParams(), &params.Email)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter email: %s", err))
	}

	// ------------- Optional query parameter "role" -------------

	err = runtime.BindQueryParameter("form", true, false, "role", ctx.QueryParams(), &params.Role)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter role: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.ListClinicians(ctx, clinicId, params)
	return err
}

// CreateClinician converts echo context to params.
func (w *ServerInterfaceWrapper) CreateClinician(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "clinicId" -------------
	var clinicId ClinicId

	err = runtime.BindStyledParameterWithLocation("simple", false, "clinicId", runtime.ParamLocationPath, ctx.Param("clinicId"), &clinicId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter clinicId: %s", err))
	}

	ctx.Set(SessionTokenScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.CreateClinician(ctx, clinicId)
	return err
}

// DeleteClinician converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteClinician(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "clinicId" -------------
	var clinicId ClinicId

	err = runtime.BindStyledParameterWithLocation("simple", false, "clinicId", runtime.ParamLocationPath, ctx.Param("clinicId"), &clinicId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter clinicId: %s", err))
	}

	// ------------- Path parameter "clinicianId" -------------
	var clinicianId ClinicianId

	err = runtime.BindStyledParameterWithLocation("simple", false, "clinicianId", runtime.ParamLocationPath, ctx.Param("clinicianId"), &clinicianId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter clinicianId: %s", err))
	}

	ctx.Set(SessionTokenScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.DeleteClinician(ctx, clinicId, clinicianId)
	return err
}

// GetClinician converts echo context to params.
func (w *ServerInterfaceWrapper) GetClinician(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "clinicId" -------------
	var clinicId ClinicId

	err = runtime.BindStyledParameterWithLocation("simple", false, "clinicId", runtime.ParamLocationPath, ctx.Param("clinicId"), &clinicId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter clinicId: %s", err))
	}

	// ------------- Path parameter "clinicianId" -------------
	var clinicianId ClinicianId

	err = runtime.BindStyledParameterWithLocation("simple", false, "clinicianId", runtime.ParamLocationPath, ctx.Param("clinicianId"), &clinicianId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter clinicianId: %s", err))
	}

	ctx.Set(SessionTokenScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetClinician(ctx, clinicId, clinicianId)
	return err
}

// UpdateClinician converts echo context to params.
func (w *ServerInterfaceWrapper) UpdateClinician(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "clinicId" -------------
	var clinicId ClinicId

	err = runtime.BindStyledParameterWithLocation("simple", false, "clinicId", runtime.ParamLocationPath, ctx.Param("clinicId"), &clinicId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter clinicId: %s", err))
	}

	// ------------- Path parameter "clinicianId" -------------
	var clinicianId ClinicianId

	err = runtime.BindStyledParameterWithLocation("simple", false, "clinicianId", runtime.ParamLocationPath, ctx.Param("clinicianId"), &clinicianId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter clinicianId: %s", err))
	}

	ctx.Set(SessionTokenScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.UpdateClinician(ctx, clinicId, clinicianId)
	return err
}

// DeleteInvitedClinician converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteInvitedClinician(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "clinicId" -------------
	var clinicId ClinicId

	err = runtime.BindStyledParameterWithLocation("simple", false, "clinicId", runtime.ParamLocationPath, ctx.Param("clinicId"), &clinicId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter clinicId: %s", err))
	}

	// ------------- Path parameter "inviteId" -------------
	var inviteId InviteId

	err = runtime.BindStyledParameterWithLocation("simple", false, "inviteId", runtime.ParamLocationPath, ctx.Param("inviteId"), &inviteId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter inviteId: %s", err))
	}

	ctx.Set(SessionTokenScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.DeleteInvitedClinician(ctx, clinicId, inviteId)
	return err
}

// GetInvitedClinician converts echo context to params.
func (w *ServerInterfaceWrapper) GetInvitedClinician(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "clinicId" -------------
	var clinicId ClinicId

	err = runtime.BindStyledParameterWithLocation("simple", false, "clinicId", runtime.ParamLocationPath, ctx.Param("clinicId"), &clinicId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter clinicId: %s", err))
	}

	// ------------- Path parameter "inviteId" -------------
	var inviteId InviteId

	err = runtime.BindStyledParameterWithLocation("simple", false, "inviteId", runtime.ParamLocationPath, ctx.Param("inviteId"), &inviteId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter inviteId: %s", err))
	}

	ctx.Set(SessionTokenScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetInvitedClinician(ctx, clinicId, inviteId)
	return err
}

// AssociateClinicianToUser converts echo context to params.
func (w *ServerInterfaceWrapper) AssociateClinicianToUser(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "clinicId" -------------
	var clinicId ClinicId

	err = runtime.BindStyledParameterWithLocation("simple", false, "clinicId", runtime.ParamLocationPath, ctx.Param("clinicId"), &clinicId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter clinicId: %s", err))
	}

	// ------------- Path parameter "inviteId" -------------
	var inviteId InviteId

	err = runtime.BindStyledParameterWithLocation("simple", false, "inviteId", runtime.ParamLocationPath, ctx.Param("inviteId"), &inviteId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter inviteId: %s", err))
	}

	ctx.Set(SessionTokenScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.AssociateClinicianToUser(ctx, clinicId, inviteId)
	return err
}

// TriggerInitialMigration converts echo context to params.
func (w *ServerInterfaceWrapper) TriggerInitialMigration(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "clinicId" -------------
	var clinicId string

	err = runtime.BindStyledParameterWithLocation("simple", false, "clinicId", runtime.ParamLocationPath, ctx.Param("clinicId"), &clinicId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter clinicId: %s", err))
	}

	ctx.Set(SessionTokenScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.TriggerInitialMigration(ctx, clinicId)
	return err
}

// ListMigrations converts echo context to params.
func (w *ServerInterfaceWrapper) ListMigrations(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "clinicId" -------------
	var clinicId string

	err = runtime.BindStyledParameterWithLocation("simple", false, "clinicId", runtime.ParamLocationPath, ctx.Param("clinicId"), &clinicId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter clinicId: %s", err))
	}

	ctx.Set(SessionTokenScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.ListMigrations(ctx, clinicId)
	return err
}

// MigrateLegacyClinicianPatients converts echo context to params.
func (w *ServerInterfaceWrapper) MigrateLegacyClinicianPatients(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "clinicId" -------------
	var clinicId string

	err = runtime.BindStyledParameterWithLocation("simple", false, "clinicId", runtime.ParamLocationPath, ctx.Param("clinicId"), &clinicId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter clinicId: %s", err))
	}

	ctx.Set(SessionTokenScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.MigrateLegacyClinicianPatients(ctx, clinicId)
	return err
}

// GetMigration converts echo context to params.
func (w *ServerInterfaceWrapper) GetMigration(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "clinicId" -------------
	var clinicId Id

	err = runtime.BindStyledParameterWithLocation("simple", false, "clinicId", runtime.ParamLocationPath, ctx.Param("clinicId"), &clinicId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter clinicId: %s", err))
	}

	// ------------- Path parameter "userId" -------------
	var userId UserId

	err = runtime.BindStyledParameterWithLocation("simple", false, "userId", runtime.ParamLocationPath, ctx.Param("userId"), &userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter userId: %s", err))
	}

	ctx.Set(SessionTokenScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetMigration(ctx, clinicId, userId)
	return err
}

// UpdateMigration converts echo context to params.
func (w *ServerInterfaceWrapper) UpdateMigration(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "clinicId" -------------
	var clinicId Id

	err = runtime.BindStyledParameterWithLocation("simple", false, "clinicId", runtime.ParamLocationPath, ctx.Param("clinicId"), &clinicId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter clinicId: %s", err))
	}

	// ------------- Path parameter "userId" -------------
	var userId UserId

	err = runtime.BindStyledParameterWithLocation("simple", false, "userId", runtime.ParamLocationPath, ctx.Param("userId"), &userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter userId: %s", err))
	}

	ctx.Set(SessionTokenScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.UpdateMigration(ctx, clinicId, userId)
	return err
}

// ListPatients converts echo context to params.
func (w *ServerInterfaceWrapper) ListPatients(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "clinicId" -------------
	var clinicId ClinicId

	err = runtime.BindStyledParameterWithLocation("simple", false, "clinicId", runtime.ParamLocationPath, ctx.Param("clinicId"), &clinicId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter clinicId: %s", err))
	}

	ctx.Set(SessionTokenScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params ListPatientsParams
	// ------------- Optional query parameter "search" -------------

	err = runtime.BindQueryParameter("form", true, false, "search", ctx.QueryParams(), &params.Search)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter search: %s", err))
	}

	// ------------- Optional query parameter "offset" -------------

	err = runtime.BindQueryParameter("form", true, false, "offset", ctx.QueryParams(), &params.Offset)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter offset: %s", err))
	}

	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", ctx.QueryParams(), &params.Limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter limit: %s", err))
	}

	// ------------- Optional query parameter "sort" -------------

	err = runtime.BindQueryParameter("form", true, false, "sort", ctx.QueryParams(), &params.Sort)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter sort: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.ListPatients(ctx, clinicId, params)
	return err
}

// CreatePatientAccount converts echo context to params.
func (w *ServerInterfaceWrapper) CreatePatientAccount(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "clinicId" -------------
	var clinicId ClinicId

	err = runtime.BindStyledParameterWithLocation("simple", false, "clinicId", runtime.ParamLocationPath, ctx.Param("clinicId"), &clinicId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter clinicId: %s", err))
	}

	ctx.Set(SessionTokenScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.CreatePatientAccount(ctx, clinicId)
	return err
}

// DeletePatient converts echo context to params.
func (w *ServerInterfaceWrapper) DeletePatient(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "clinicId" -------------
	var clinicId ClinicId

	err = runtime.BindStyledParameterWithLocation("simple", false, "clinicId", runtime.ParamLocationPath, ctx.Param("clinicId"), &clinicId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter clinicId: %s", err))
	}

	// ------------- Path parameter "patientId" -------------
	var patientId PatientId

	err = runtime.BindStyledParameterWithLocation("simple", false, "patientId", runtime.ParamLocationPath, ctx.Param("patientId"), &patientId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter patientId: %s", err))
	}

	ctx.Set(SessionTokenScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.DeletePatient(ctx, clinicId, patientId)
	return err
}

// GetPatient converts echo context to params.
func (w *ServerInterfaceWrapper) GetPatient(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "clinicId" -------------
	var clinicId ClinicId

	err = runtime.BindStyledParameterWithLocation("simple", false, "clinicId", runtime.ParamLocationPath, ctx.Param("clinicId"), &clinicId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter clinicId: %s", err))
	}

	// ------------- Path parameter "patientId" -------------
	var patientId PatientId

	err = runtime.BindStyledParameterWithLocation("simple", false, "patientId", runtime.ParamLocationPath, ctx.Param("patientId"), &patientId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter patientId: %s", err))
	}

	ctx.Set(SessionTokenScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetPatient(ctx, clinicId, patientId)
	return err
}

// CreatePatientFromUser converts echo context to params.
func (w *ServerInterfaceWrapper) CreatePatientFromUser(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "clinicId" -------------
	var clinicId ClinicId

	err = runtime.BindStyledParameterWithLocation("simple", false, "clinicId", runtime.ParamLocationPath, ctx.Param("clinicId"), &clinicId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter clinicId: %s", err))
	}

	// ------------- Path parameter "patientId" -------------
	var patientId PatientId

	err = runtime.BindStyledParameterWithLocation("simple", false, "patientId", runtime.ParamLocationPath, ctx.Param("patientId"), &patientId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter patientId: %s", err))
	}

	ctx.Set(SessionTokenScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.CreatePatientFromUser(ctx, clinicId, patientId)
	return err
}

// UpdatePatient converts echo context to params.
func (w *ServerInterfaceWrapper) UpdatePatient(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "clinicId" -------------
	var clinicId ClinicId

	err = runtime.BindStyledParameterWithLocation("simple", false, "clinicId", runtime.ParamLocationPath, ctx.Param("clinicId"), &clinicId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter clinicId: %s", err))
	}

	// ------------- Path parameter "patientId" -------------
	var patientId PatientId

	err = runtime.BindStyledParameterWithLocation("simple", false, "patientId", runtime.ParamLocationPath, ctx.Param("patientId"), &patientId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter patientId: %s", err))
	}

	ctx.Set(SessionTokenScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.UpdatePatient(ctx, clinicId, patientId)
	return err
}

// UpdatePatientPermissions converts echo context to params.
func (w *ServerInterfaceWrapper) UpdatePatientPermissions(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "clinicId" -------------
	var clinicId ClinicId

	err = runtime.BindStyledParameterWithLocation("simple", false, "clinicId", runtime.ParamLocationPath, ctx.Param("clinicId"), &clinicId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter clinicId: %s", err))
	}

	// ------------- Path parameter "patientId" -------------
	var patientId PatientId

	err = runtime.BindStyledParameterWithLocation("simple", false, "patientId", runtime.ParamLocationPath, ctx.Param("patientId"), &patientId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter patientId: %s", err))
	}

	ctx.Set(SessionTokenScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.UpdatePatientPermissions(ctx, clinicId, patientId)
	return err
}

// DeletePatientPermission converts echo context to params.
func (w *ServerInterfaceWrapper) DeletePatientPermission(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "clinicId" -------------
	var clinicId ClinicId

	err = runtime.BindStyledParameterWithLocation("simple", false, "clinicId", runtime.ParamLocationPath, ctx.Param("clinicId"), &clinicId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter clinicId: %s", err))
	}

	// ------------- Path parameter "patientId" -------------
	var patientId PatientId

	err = runtime.BindStyledParameterWithLocation("simple", false, "patientId", runtime.ParamLocationPath, ctx.Param("patientId"), &patientId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter patientId: %s", err))
	}

	// ------------- Path parameter "permission" -------------
	var permission string

	err = runtime.BindStyledParameterWithLocation("simple", false, "permission", runtime.ParamLocationPath, ctx.Param("permission"), &permission)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter permission: %s", err))
	}

	ctx.Set(SessionTokenScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.DeletePatientPermission(ctx, clinicId, patientId, permission)
	return err
}

// ListClinicsForPatient converts echo context to params.
func (w *ServerInterfaceWrapper) ListClinicsForPatient(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "userId" -------------
	var userId UserId

	err = runtime.BindStyledParameterWithLocation("simple", false, "userId", runtime.ParamLocationPath, ctx.Param("userId"), &userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter userId: %s", err))
	}

	ctx.Set(SessionTokenScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params ListClinicsForPatientParams
	// ------------- Optional query parameter "offset" -------------

	err = runtime.BindQueryParameter("form", true, false, "offset", ctx.QueryParams(), &params.Offset)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter offset: %s", err))
	}

	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", ctx.QueryParams(), &params.Limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter limit: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.ListClinicsForPatient(ctx, userId, params)
	return err
}

// DeleteUserFromClinics converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteUserFromClinics(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "userId" -------------
	var userId UserId

	err = runtime.BindStyledParameterWithLocation("simple", false, "userId", runtime.ParamLocationPath, ctx.Param("userId"), &userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter userId: %s", err))
	}

	ctx.Set(SessionTokenScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.DeleteUserFromClinics(ctx, userId)
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
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/v1/clinicians/:userId/clinics", wrapper.ListClinicsForClinician)
	router.POST(baseURL+"/v1/clinicians/:userId/migrate", wrapper.EnableNewClinicExperience)
	router.GET(baseURL+"/v1/clinics", wrapper.ListClinics)
	router.POST(baseURL+"/v1/clinics", wrapper.CreateClinic)
	router.GET(baseURL+"/v1/clinics/share_code/:shareCode", wrapper.GetClinicByShareCode)
	router.GET(baseURL+"/v1/clinics/:clinicId", wrapper.GetClinic)
	router.PUT(baseURL+"/v1/clinics/:clinicId", wrapper.UpdateClinic)
	router.GET(baseURL+"/v1/clinics/:clinicId/clinicians", wrapper.ListClinicians)
	router.POST(baseURL+"/v1/clinics/:clinicId/clinicians", wrapper.CreateClinician)
	router.DELETE(baseURL+"/v1/clinics/:clinicId/clinicians/:clinicianId", wrapper.DeleteClinician)
	router.GET(baseURL+"/v1/clinics/:clinicId/clinicians/:clinicianId", wrapper.GetClinician)
	router.PUT(baseURL+"/v1/clinics/:clinicId/clinicians/:clinicianId", wrapper.UpdateClinician)
	router.DELETE(baseURL+"/v1/clinics/:clinicId/invites/clinicians/:inviteId/clinician", wrapper.DeleteInvitedClinician)
	router.GET(baseURL+"/v1/clinics/:clinicId/invites/clinicians/:inviteId/clinician", wrapper.GetInvitedClinician)
	router.PATCH(baseURL+"/v1/clinics/:clinicId/invites/clinicians/:inviteId/clinician", wrapper.AssociateClinicianToUser)
	router.POST(baseURL+"/v1/clinics/:clinicId/migrate", wrapper.TriggerInitialMigration)
	router.GET(baseURL+"/v1/clinics/:clinicId/migrations", wrapper.ListMigrations)
	router.POST(baseURL+"/v1/clinics/:clinicId/migrations", wrapper.MigrateLegacyClinicianPatients)
	router.GET(baseURL+"/v1/clinics/:clinicId/migrations/:userId", wrapper.GetMigration)
	router.PATCH(baseURL+"/v1/clinics/:clinicId/migrations/:userId", wrapper.UpdateMigration)
	router.GET(baseURL+"/v1/clinics/:clinicId/patients", wrapper.ListPatients)
	router.POST(baseURL+"/v1/clinics/:clinicId/patients", wrapper.CreatePatientAccount)
	router.DELETE(baseURL+"/v1/clinics/:clinicId/patients/:patientId", wrapper.DeletePatient)
	router.GET(baseURL+"/v1/clinics/:clinicId/patients/:patientId", wrapper.GetPatient)
	router.POST(baseURL+"/v1/clinics/:clinicId/patients/:patientId", wrapper.CreatePatientFromUser)
	router.PUT(baseURL+"/v1/clinics/:clinicId/patients/:patientId", wrapper.UpdatePatient)
	router.PUT(baseURL+"/v1/clinics/:clinicId/patients/:patientId/permissions", wrapper.UpdatePatientPermissions)
	router.DELETE(baseURL+"/v1/clinics/:clinicId/patients/:patientId/permissions/:permission", wrapper.DeletePatientPermission)
	router.GET(baseURL+"/v1/patients/:userId/clinics", wrapper.ListClinicsForPatient)
	router.DELETE(baseURL+"/v1/users/:userId/clinics", wrapper.DeleteUserFromClinics)

}

