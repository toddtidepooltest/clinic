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
	// GetClinics
	// (GET /clinics)
	GetClinics(ctx echo.Context, params GetClinicsParams) error
	// AddClinic
	// (POST /clinics)
	PostClinics(ctx echo.Context) error
	// DeleteClinic
	// (DELETE /clinics/{clinicid})
	DeleteClinicsClinicid(ctx echo.Context, clinicid string) error
	// GetClinic
	// (GET /clinics/{clinicid})
	GetClinicsClinicid(ctx echo.Context, clinicid string) error
	// ModifyClinic
	// (PATCH /clinics/{clinicid})
	PatchClinicsClinicid(ctx echo.Context, clinicid string) error
	// GetCliniciansFromClinic
	// (GET /clinics/{clinicid}/clinicians)
	GetClinicsClinicidClinicians(ctx echo.Context, clinicid string, params GetClinicsClinicidCliniciansParams) error
	// AddClinicianToClinic
	// (POST /clinics/{clinicid}/clinicians)
	PostClinicsClinicidClinicians(ctx echo.Context, clinicid string) error
	// DeleteClinicianFromClinic
	// (DELETE /clinics/{clinicid}/clinicians/{clinicianid})
	DeleteClinicsClinicidCliniciansClinicianid(ctx echo.Context, clinicid string, clinicianid string) error
	// GetClinician
	// (GET /clinics/{clinicid}/clinicians/{clinicianid})
	GetClinicsClinicidCliniciansClinicianid(ctx echo.Context, clinicid string, clinicianid string) error
	// ModifyClinicClinician
	// (PATCH /clinics/{clinicid}/clinicians/{clinicianid})
	PatchClinicsClinicidCliniciansClinicianid(ctx echo.Context, clinicid string, clinicianid string) error
	// GetPatientsForClinic
	// (GET /clinics/{clinicid}/patients)
	GetClinicsClinicidPatients(ctx echo.Context, clinicid string, params GetClinicsClinicidPatientsParams) error
	// AddPatientToClinic
	// (POST /clinics/{clinicid}/patients)
	PostClinicsClinicidPatients(ctx echo.Context, clinicid string) error
	// DeletePatientFromClinic
	// (DELETE /clinics/{clinicid}/patients/{patientid})
	DeleteClinicClinicidPatientPatientid(ctx echo.Context, clinicid string, patientid string) error
	// GetPatientFromClinic
	// (GET /clinics/{clinicid}/patients/{patientid})
	GetClinicsClinicidPatientPatientid(ctx echo.Context, clinicid string, patientid string) error
	// ModifyClinicPatient
	// (PATCH /clinics/{clinicid}/patients/{patientid})
	PatchClinicsClinicidPatientsPatientid(ctx echo.Context, clinicid string, patientid string) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetClinics converts echo context to params.
func (w *ServerInterfaceWrapper) GetClinics(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetClinicsParams
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

	// ------------- Optional query parameter "sortOrder" -------------

	err = runtime.BindQueryParameter("form", true, false, "sortOrder", ctx.QueryParams(), &params.SortOrder)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter sortOrder: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetClinics(ctx, params)
	return err
}

// PostClinics converts echo context to params.
func (w *ServerInterfaceWrapper) PostClinics(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PostClinics(ctx)
	return err
}

// DeleteClinicsClinicid converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteClinicsClinicid(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "clinicid" -------------
	var clinicid string

	err = runtime.BindStyledParameter("simple", false, "clinicid", ctx.Param("clinicid"), &clinicid)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter clinicid: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.DeleteClinicsClinicid(ctx, clinicid)
	return err
}

// GetClinicsClinicid converts echo context to params.
func (w *ServerInterfaceWrapper) GetClinicsClinicid(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "clinicid" -------------
	var clinicid string

	err = runtime.BindStyledParameter("simple", false, "clinicid", ctx.Param("clinicid"), &clinicid)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter clinicid: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetClinicsClinicid(ctx, clinicid)
	return err
}

// PatchClinicsClinicid converts echo context to params.
func (w *ServerInterfaceWrapper) PatchClinicsClinicid(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "clinicid" -------------
	var clinicid string

	err = runtime.BindStyledParameter("simple", false, "clinicid", ctx.Param("clinicid"), &clinicid)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter clinicid: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PatchClinicsClinicid(ctx, clinicid)
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

	// Parameter object where we will unmarshal all parameters from the context
	var params GetClinicsClinicidCliniciansParams
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

	// ------------- Optional query parameter "sortOrder" -------------

	err = runtime.BindQueryParameter("form", true, false, "sortOrder", ctx.QueryParams(), &params.SortOrder)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter sortOrder: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetClinicsClinicidClinicians(ctx, clinicid, params)
	return err
}

// PostClinicsClinicidClinicians converts echo context to params.
func (w *ServerInterfaceWrapper) PostClinicsClinicidClinicians(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "clinicid" -------------
	var clinicid string

	err = runtime.BindStyledParameter("simple", false, "clinicid", ctx.Param("clinicid"), &clinicid)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter clinicid: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PostClinicsClinicidClinicians(ctx, clinicid)
	return err
}

// DeleteClinicsClinicidCliniciansClinicianid converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteClinicsClinicidCliniciansClinicianid(ctx echo.Context) error {
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
	err = w.Handler.DeleteClinicsClinicidCliniciansClinicianid(ctx, clinicid, clinicianid)
	return err
}

// GetClinicsClinicidCliniciansClinicianid converts echo context to params.
func (w *ServerInterfaceWrapper) GetClinicsClinicidCliniciansClinicianid(ctx echo.Context) error {
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
	err = w.Handler.GetClinicsClinicidCliniciansClinicianid(ctx, clinicid, clinicianid)
	return err
}

// PatchClinicsClinicidCliniciansClinicianid converts echo context to params.
func (w *ServerInterfaceWrapper) PatchClinicsClinicidCliniciansClinicianid(ctx echo.Context) error {
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
	err = w.Handler.PatchClinicsClinicidCliniciansClinicianid(ctx, clinicid, clinicianid)
	return err
}

// GetClinicsClinicidPatients converts echo context to params.
func (w *ServerInterfaceWrapper) GetClinicsClinicidPatients(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "clinicid" -------------
	var clinicid string

	err = runtime.BindStyledParameter("simple", false, "clinicid", ctx.Param("clinicid"), &clinicid)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter clinicid: %s", err))
	}

	// Parameter object where we will unmarshal all parameters from the context
	var params GetClinicsClinicidPatientsParams
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

	// ------------- Optional query parameter "sortOrder" -------------

	err = runtime.BindQueryParameter("form", true, false, "sortOrder", ctx.QueryParams(), &params.SortOrder)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter sortOrder: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetClinicsClinicidPatients(ctx, clinicid, params)
	return err
}

// PostClinicsClinicidPatients converts echo context to params.
func (w *ServerInterfaceWrapper) PostClinicsClinicidPatients(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "clinicid" -------------
	var clinicid string

	err = runtime.BindStyledParameter("simple", false, "clinicid", ctx.Param("clinicid"), &clinicid)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter clinicid: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PostClinicsClinicidPatients(ctx, clinicid)
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

// GetClinicsClinicidPatientPatientid converts echo context to params.
func (w *ServerInterfaceWrapper) GetClinicsClinicidPatientPatientid(ctx echo.Context) error {
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
	err = w.Handler.GetClinicsClinicidPatientPatientid(ctx, clinicid, patientid)
	return err
}

// PatchClinicsClinicidPatientsPatientid converts echo context to params.
func (w *ServerInterfaceWrapper) PatchClinicsClinicidPatientsPatientid(ctx echo.Context) error {
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
	err = w.Handler.PatchClinicsClinicidPatientsPatientid(ctx, clinicid, patientid)
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

	router.GET("/clinics", wrapper.GetClinics)
	router.POST("/clinics", wrapper.PostClinics)
	router.DELETE("/clinics/:clinicid", wrapper.DeleteClinicsClinicid)
	router.GET("/clinics/:clinicid", wrapper.GetClinicsClinicid)
	router.PATCH("/clinics/:clinicid", wrapper.PatchClinicsClinicid)
	router.GET("/clinics/:clinicid/clinicians", wrapper.GetClinicsClinicidClinicians)
	router.POST("/clinics/:clinicid/clinicians", wrapper.PostClinicsClinicidClinicians)
	router.DELETE("/clinics/:clinicid/clinicians/:clinicianid", wrapper.DeleteClinicsClinicidCliniciansClinicianid)
	router.GET("/clinics/:clinicid/clinicians/:clinicianid", wrapper.GetClinicsClinicidCliniciansClinicianid)
	router.PATCH("/clinics/:clinicid/clinicians/:clinicianid", wrapper.PatchClinicsClinicidCliniciansClinicianid)
	router.GET("/clinics/:clinicid/patients", wrapper.GetClinicsClinicidPatients)
	router.POST("/clinics/:clinicid/patients", wrapper.PostClinicsClinicidPatients)
	router.DELETE("/clinics/:clinicid/patients/:patientid", wrapper.DeleteClinicClinicidPatientPatientid)
	router.GET("/clinics/:clinicid/patients/:patientid", wrapper.GetClinicsClinicidPatientPatientid)
	router.PATCH("/clinics/:clinicid/patients/:patientid", wrapper.PatchClinicsClinicidPatientsPatientid)

}

