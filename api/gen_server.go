// Package Clinic provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package api

import (
	"fmt"
	"net/http"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/labstack/echo/v4"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// List Clinics
	// (GET /v1/clinics)
	ListClinics(ctx echo.Context, params ListClinicsParams) error
	// Create Clinic
	// (POST /v1/clinics)
	CreateClinic(ctx echo.Context) error
	// Get Clinic
	// (GET /v1/clinics/{clinicId})
	GetClinic(ctx echo.Context, clinicId string) error
	// Update Clinic
	// (PATCH /v1/clinics/{clinicId})
	UpdateClinic(ctx echo.Context, clinicId string) error
	// List Clinicians
	// (GET /v1/clinics/{clinicId}/clinicians)
	ListClinicians(ctx echo.Context, clinicId string, params ListCliniciansParams) error
	// Delete Clinician
	// (DELETE /v1/clinics/{clinicId}/clinicians/{clinicianId})
	DeleteClinician(ctx echo.Context, clinicId string, clinicianId string) error
	// Get Clinician
	// (GET /v1/clinics/{clinicId}/clinicians/{clinicianId})
	GetClinician(ctx echo.Context, clinicId string, clinicianId string) error
	// Update Clinician
	// (PATCH /v1/clinics/{clinicId}/clinicians/{clinicianId})
	UpdateClinician(ctx echo.Context, clinicId string, clinicianId string) error
	// Invite Clinician
	// (POST /v1/clinics/{clinicId}/invites)
	InviteClinician(ctx echo.Context, clinicId string) error

	// (DELETE /v1/clinics/{clinicId}/invites/{inviteId})
	DeleteInvite(ctx echo.Context, clinicId string, inviteId string) error

	// (PATCH /v1/clinics/{clinicId}/invites/{inviteId})
	ResendInvite(ctx echo.Context, clinicId string, inviteId string) error
	// Accept Invite
	// (PUT /v1/clinics/{clinicId}/invites/{inviteId})
	AcceptInvite(ctx echo.Context, clinicId string, inviteId string) error
	// List Patients
	// (GET /v1/clinics/{clinicId}/patients)
	ListPatients(ctx echo.Context, clinicId string, params ListPatientsParams) error
	// Create Patient Account
	// (POST /v1/clinics/{clinicId}/patients)
	CreatePatientAccount(ctx echo.Context, clinicId string) error
	// Get Patient
	// (GET /v1/clinics/{clinicId}/patients/{patientId})
	GetPatient(ctx echo.Context, clinicId string, patientId string) error
	// Create Patient from Existing User
	// (POST /v1/clinics/{clinicId}/patients/{patientId})
	CreatePatientFromUser(ctx echo.Context, clinicId string, patientId string) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// ListClinics converts echo context to params.
func (w *ServerInterfaceWrapper) ListClinics(ctx echo.Context) error {
	var err error

	ctx.Set(SessionTokenScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params ListClinicsParams
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

	// ------------- Optional query parameter "clinicianId" -------------

	err = runtime.BindQueryParameter("form", true, false, "clinicianId", ctx.QueryParams(), &params.ClinicianId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter clinicianId: %s", err))
	}

	// ------------- Optional query parameter "patientId" -------------

	err = runtime.BindQueryParameter("form", true, false, "patientId", ctx.QueryParams(), &params.PatientId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter patientId: %s", err))
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

// GetClinic converts echo context to params.
func (w *ServerInterfaceWrapper) GetClinic(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "clinicId" -------------
	var clinicId string

	err = runtime.BindStyledParameter("simple", false, "clinicId", ctx.Param("clinicId"), &clinicId)
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
	var clinicId string

	err = runtime.BindStyledParameter("simple", false, "clinicId", ctx.Param("clinicId"), &clinicId)
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
	var clinicId string

	err = runtime.BindStyledParameter("simple", false, "clinicId", ctx.Param("clinicId"), &clinicId)
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

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.ListClinicians(ctx, clinicId, params)
	return err
}

// DeleteClinician converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteClinician(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "clinicId" -------------
	var clinicId string

	err = runtime.BindStyledParameter("simple", false, "clinicId", ctx.Param("clinicId"), &clinicId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter clinicId: %s", err))
	}

	// ------------- Path parameter "clinicianId" -------------
	var clinicianId string

	err = runtime.BindStyledParameter("simple", false, "clinicianId", ctx.Param("clinicianId"), &clinicianId)
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
	var clinicId string

	err = runtime.BindStyledParameter("simple", false, "clinicId", ctx.Param("clinicId"), &clinicId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter clinicId: %s", err))
	}

	// ------------- Path parameter "clinicianId" -------------
	var clinicianId string

	err = runtime.BindStyledParameter("simple", false, "clinicianId", ctx.Param("clinicianId"), &clinicianId)
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
	var clinicId string

	err = runtime.BindStyledParameter("simple", false, "clinicId", ctx.Param("clinicId"), &clinicId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter clinicId: %s", err))
	}

	// ------------- Path parameter "clinicianId" -------------
	var clinicianId string

	err = runtime.BindStyledParameter("simple", false, "clinicianId", ctx.Param("clinicianId"), &clinicianId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter clinicianId: %s", err))
	}

	ctx.Set(SessionTokenScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.UpdateClinician(ctx, clinicId, clinicianId)
	return err
}

// InviteClinician converts echo context to params.
func (w *ServerInterfaceWrapper) InviteClinician(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "clinicId" -------------
	var clinicId string

	err = runtime.BindStyledParameter("simple", false, "clinicId", ctx.Param("clinicId"), &clinicId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter clinicId: %s", err))
	}

	ctx.Set(SessionTokenScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.InviteClinician(ctx, clinicId)
	return err
}

// DeleteInvite converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteInvite(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "clinicId" -------------
	var clinicId string

	err = runtime.BindStyledParameter("simple", false, "clinicId", ctx.Param("clinicId"), &clinicId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter clinicId: %s", err))
	}

	// ------------- Path parameter "inviteId" -------------
	var inviteId string

	err = runtime.BindStyledParameter("simple", false, "inviteId", ctx.Param("inviteId"), &inviteId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter inviteId: %s", err))
	}

	ctx.Set(SessionTokenScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.DeleteInvite(ctx, clinicId, inviteId)
	return err
}

// ResendInvite converts echo context to params.
func (w *ServerInterfaceWrapper) ResendInvite(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "clinicId" -------------
	var clinicId string

	err = runtime.BindStyledParameter("simple", false, "clinicId", ctx.Param("clinicId"), &clinicId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter clinicId: %s", err))
	}

	// ------------- Path parameter "inviteId" -------------
	var inviteId string

	err = runtime.BindStyledParameter("simple", false, "inviteId", ctx.Param("inviteId"), &inviteId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter inviteId: %s", err))
	}

	ctx.Set(SessionTokenScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.ResendInvite(ctx, clinicId, inviteId)
	return err
}

// AcceptInvite converts echo context to params.
func (w *ServerInterfaceWrapper) AcceptInvite(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "clinicId" -------------
	var clinicId string

	err = runtime.BindStyledParameter("simple", false, "clinicId", ctx.Param("clinicId"), &clinicId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter clinicId: %s", err))
	}

	// ------------- Path parameter "inviteId" -------------
	var inviteId string

	err = runtime.BindStyledParameter("simple", false, "inviteId", ctx.Param("inviteId"), &inviteId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter inviteId: %s", err))
	}

	ctx.Set(SessionTokenScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.AcceptInvite(ctx, clinicId, inviteId)
	return err
}

// ListPatients converts echo context to params.
func (w *ServerInterfaceWrapper) ListPatients(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "clinicId" -------------
	var clinicId string

	err = runtime.BindStyledParameter("simple", false, "clinicId", ctx.Param("clinicId"), &clinicId)
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

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.ListPatients(ctx, clinicId, params)
	return err
}

// CreatePatientAccount converts echo context to params.
func (w *ServerInterfaceWrapper) CreatePatientAccount(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "clinicId" -------------
	var clinicId string

	err = runtime.BindStyledParameter("simple", false, "clinicId", ctx.Param("clinicId"), &clinicId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter clinicId: %s", err))
	}

	ctx.Set(SessionTokenScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.CreatePatientAccount(ctx, clinicId)
	return err
}

// GetPatient converts echo context to params.
func (w *ServerInterfaceWrapper) GetPatient(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "clinicId" -------------
	var clinicId string

	err = runtime.BindStyledParameter("simple", false, "clinicId", ctx.Param("clinicId"), &clinicId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter clinicId: %s", err))
	}

	// ------------- Path parameter "patientId" -------------
	var patientId string

	err = runtime.BindStyledParameter("simple", false, "patientId", ctx.Param("patientId"), &patientId)
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
	var clinicId string

	err = runtime.BindStyledParameter("simple", false, "clinicId", ctx.Param("clinicId"), &clinicId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter clinicId: %s", err))
	}

	// ------------- Path parameter "patientId" -------------
	var patientId string

	err = runtime.BindStyledParameter("simple", false, "patientId", ctx.Param("patientId"), &patientId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter patientId: %s", err))
	}

	ctx.Set(SessionTokenScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.CreatePatientFromUser(ctx, clinicId, patientId)
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

	router.GET(baseURL+"/v1/clinics", wrapper.ListClinics)
	router.POST(baseURL+"/v1/clinics", wrapper.CreateClinic)
	router.GET(baseURL+"/v1/clinics/:clinicId", wrapper.GetClinic)
	router.PATCH(baseURL+"/v1/clinics/:clinicId", wrapper.UpdateClinic)
	router.GET(baseURL+"/v1/clinics/:clinicId/clinicians", wrapper.ListClinicians)
	router.DELETE(baseURL+"/v1/clinics/:clinicId/clinicians/:clinicianId", wrapper.DeleteClinician)
	router.GET(baseURL+"/v1/clinics/:clinicId/clinicians/:clinicianId", wrapper.GetClinician)
	router.PATCH(baseURL+"/v1/clinics/:clinicId/clinicians/:clinicianId", wrapper.UpdateClinician)
	router.POST(baseURL+"/v1/clinics/:clinicId/invites", wrapper.InviteClinician)
	router.DELETE(baseURL+"/v1/clinics/:clinicId/invites/:inviteId", wrapper.DeleteInvite)
	router.PATCH(baseURL+"/v1/clinics/:clinicId/invites/:inviteId", wrapper.ResendInvite)
	router.PUT(baseURL+"/v1/clinics/:clinicId/invites/:inviteId", wrapper.AcceptInvite)
	router.GET(baseURL+"/v1/clinics/:clinicId/patients", wrapper.ListPatients)
	router.POST(baseURL+"/v1/clinics/:clinicId/patients", wrapper.CreatePatientAccount)
	router.GET(baseURL+"/v1/clinics/:clinicId/patients/:patientId", wrapper.GetPatient)
	router.POST(baseURL+"/v1/clinics/:clinicId/patients/:patientId", wrapper.CreatePatientFromUser)

}

