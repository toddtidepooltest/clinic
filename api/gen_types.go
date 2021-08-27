// Package api provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package api

import (
	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
)

const (
	SessionTokenScopes = "sessionToken.Scopes"
)

// AssociateClinicianToUser defines model for AssociateClinicianToUser.
type AssociateClinicianToUser struct {
	UserId string `json:"userId"`
}

// Clinic defines model for Clinic.
type Clinic struct {
	Address    *string `json:"address,omitempty"`
	CanMigrate bool    `json:"canMigrate"`
	City       *string `json:"city,omitempty"`
	ClinicSize *string `json:"clinicSize,omitempty"`
	ClinicType *string `json:"clinicType,omitempty"`
	Country    *string `json:"country,omitempty"`

	// String representation of a resource id
	Id           Id             `json:"id"`
	Name         string         `json:"name"`
	PhoneNumbers *[]PhoneNumber `json:"phoneNumbers,omitempty"`
	PostalCode   *string        `json:"postalCode,omitempty"`
	ShareCode    string         `json:"shareCode"`
	State        *string        `json:"state,omitempty"`
	Website      *string        `json:"website,omitempty"`
}

// Clinician defines model for Clinician.
type Clinician struct {

	// The email of the clinician
	Email string `json:"email"`

	// Tidepool User ID. Old style IDs are 10-digit strings consisting of only hexadeximcal digits. New style IDs are 36-digit [UUID v4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)).
	Id *TidepoolUserId `json:"id,omitempty"`

	// The id of the invite if it hasn't been accepted
	InviteId *string `json:"inviteId,omitempty"`

	// The name of the clinician
	Name  *string        `json:"name,omitempty"`
	Roles ClinicianRoles `json:"roles"`
}

// ClinicianClinicRelationship defines model for ClinicianClinicRelationship.
type ClinicianClinicRelationship struct {
	Clinic    Clinic    `json:"clinic"`
	Clinician Clinician `json:"clinician"`
}

// ClinicianClinicRelationships defines model for ClinicianClinicRelationships.
type ClinicianClinicRelationships []ClinicianClinicRelationship

// ClinicianRoles defines model for ClinicianRoles.
type ClinicianRoles []string

// Clinicians defines model for Clinicians.
type Clinicians []Clinician

// Clinics defines model for Clinics.
type Clinics []Clinic

// CreatePatient defines model for CreatePatient.
type CreatePatient struct {
	IsMigrated  *bool               `json:"isMigrated,omitempty"`
	Permissions *PatientPermissions `json:"permissions,omitempty"`
}

// Error defines model for Error.
type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Id defines model for Id.
type Id string

// Migration defines model for Migration.
type Migration struct {
	CreatedTime openapi_types.Date `json:"createdTime"`

	// The user id of the legacy clinician account that needs to be migrated.
	UserId string `json:"userId"`
}

// Migrations defines model for Migrations.
type Migrations []Migration

// Patient defines model for Patient.
type Patient struct {

	// YYYY-MM-DD
	BirthDate openapi_types.Date `json:"birthDate"`

	// The email of the patient
	Email *string `json:"email,omitempty"`

	// The full name of the patient
	FullName string `json:"fullName"`

	// Tidepool User ID. Old style IDs are 10-digit strings consisting of only hexadeximcal digits. New style IDs are 36-digit [UUID v4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)).
	Id TidepoolUserId `json:"id"`

	// The medical record number of the patient
	Mrn           *string             `json:"mrn,omitempty"`
	Permissions   *PatientPermissions `json:"permissions,omitempty"`
	TargetDevices *[]string           `json:"targetDevices,omitempty"`
}

// PatientClinicRelationship defines model for PatientClinicRelationship.
type PatientClinicRelationship struct {
	Clinic  Clinic  `json:"clinic"`
	Patient Patient `json:"patient"`
}

// PatientClinicRelationships defines model for PatientClinicRelationships.
type PatientClinicRelationships []PatientClinicRelationship

// PatientPermissions defines model for PatientPermissions.
type PatientPermissions struct {
	Custodian *map[string]interface{} `json:"custodian,omitempty"`
	Note      *map[string]interface{} `json:"note,omitempty"`
	Upload    *map[string]interface{} `json:"upload,omitempty"`
	View      *map[string]interface{} `json:"view,omitempty"`
}

// Patients defines model for Patients.
type Patients []Patient

// PhoneNumber defines model for PhoneNumber.
type PhoneNumber struct {
	Number string  `json:"number"`
	Type   *string `json:"type,omitempty"`
}

// TidepoolUserId defines model for TidepoolUserId.
type TidepoolUserId string

// ClinicId defines model for clinicId.
type ClinicId string

// ClinicianId defines model for clinicianId.
type ClinicianId string

// Email defines model for email.
type Email openapi_types.Email

// InviteId defines model for inviteId.
type InviteId string

// Limit defines model for limit.
type Limit int

// Offset defines model for offset.
type Offset int

// PatientId defines model for patientId.
type PatientId string

// Search defines model for search.
type Search string

// ShareCode defines model for shareCode.
type ShareCode string

// UserId defines model for userId.
type UserId TidepoolUserId

// ListClinicsForClinicianParams defines parameters for ListClinicsForClinician.
type ListClinicsForClinicianParams struct {
	Offset *Offset `json:"offset,omitempty"`
	Limit  *Limit  `json:"limit,omitempty"`
}

// ListClinicsParams defines parameters for ListClinics.
type ListClinicsParams struct {
	Limit     *Limit     `json:"limit,omitempty"`
	Offset    *Offset    `json:"offset,omitempty"`
	ShareCode *ShareCode `json:"shareCode,omitempty"`
}

// CreateClinicJSONBody defines parameters for CreateClinic.
type CreateClinicJSONBody Clinic

// UpdateClinicJSONBody defines parameters for UpdateClinic.
type UpdateClinicJSONBody Clinic

// ListCliniciansParams defines parameters for ListClinicians.
type ListCliniciansParams struct {

	// Full text search query
	Search *Search `json:"search,omitempty"`
	Offset *Offset `json:"offset,omitempty"`
	Limit  *Limit  `json:"limit,omitempty"`
	Email  *Email  `json:"email,omitempty"`
}

// CreateClinicianJSONBody defines parameters for CreateClinician.
type CreateClinicianJSONBody Clinician

// UpdateClinicianJSONBody defines parameters for UpdateClinician.
type UpdateClinicianJSONBody Clinician

// AssociateClinicianToUserJSONBody defines parameters for AssociateClinicianToUser.
type AssociateClinicianToUserJSONBody AssociateClinicianToUser

// MigrateLegacyClinicianPatientsJSONBody defines parameters for MigrateLegacyClinicianPatients.
type MigrateLegacyClinicianPatientsJSONBody Migration

// ListPatientsParams defines parameters for ListPatients.
type ListPatientsParams struct {

	// Full text search query
	Search *Search `json:"search,omitempty"`
	Offset *Offset `json:"offset,omitempty"`
	Limit  *Limit  `json:"limit,omitempty"`
}

// CreatePatientAccountJSONBody defines parameters for CreatePatientAccount.
type CreatePatientAccountJSONBody Patient

// CreatePatientFromUserJSONBody defines parameters for CreatePatientFromUser.
type CreatePatientFromUserJSONBody CreatePatient

// UpdatePatientJSONBody defines parameters for UpdatePatient.
type UpdatePatientJSONBody Patient

// UpdatePatientPermissionsJSONBody defines parameters for UpdatePatientPermissions.
type UpdatePatientPermissionsJSONBody PatientPermissions

// ListClinicsForPatientParams defines parameters for ListClinicsForPatient.
type ListClinicsForPatientParams struct {
	Offset *Offset `json:"offset,omitempty"`
	Limit  *Limit  `json:"limit,omitempty"`
}

// CreateClinicJSONRequestBody defines body for CreateClinic for application/json ContentType.
type CreateClinicJSONRequestBody CreateClinicJSONBody

// UpdateClinicJSONRequestBody defines body for UpdateClinic for application/json ContentType.
type UpdateClinicJSONRequestBody UpdateClinicJSONBody

// CreateClinicianJSONRequestBody defines body for CreateClinician for application/json ContentType.
type CreateClinicianJSONRequestBody CreateClinicianJSONBody

// UpdateClinicianJSONRequestBody defines body for UpdateClinician for application/json ContentType.
type UpdateClinicianJSONRequestBody UpdateClinicianJSONBody

// AssociateClinicianToUserJSONRequestBody defines body for AssociateClinicianToUser for application/json ContentType.
type AssociateClinicianToUserJSONRequestBody AssociateClinicianToUserJSONBody

// MigrateLegacyClinicianPatientsJSONRequestBody defines body for MigrateLegacyClinicianPatients for application/json ContentType.
type MigrateLegacyClinicianPatientsJSONRequestBody MigrateLegacyClinicianPatientsJSONBody

// CreatePatientAccountJSONRequestBody defines body for CreatePatientAccount for application/json ContentType.
type CreatePatientAccountJSONRequestBody CreatePatientAccountJSONBody

// CreatePatientFromUserJSONRequestBody defines body for CreatePatientFromUser for application/json ContentType.
type CreatePatientFromUserJSONRequestBody CreatePatientFromUserJSONBody

// UpdatePatientJSONRequestBody defines body for UpdatePatient for application/json ContentType.
type UpdatePatientJSONRequestBody UpdatePatientJSONBody

// UpdatePatientPermissionsJSONRequestBody defines body for UpdatePatientPermissions for application/json ContentType.
type UpdatePatientPermissionsJSONRequestBody UpdatePatientPermissionsJSONBody

