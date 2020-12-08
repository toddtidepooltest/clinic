// Package Clinic provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package api

// Clinic defines model for Clinic.
type Clinic struct {
	// Embedded fields due to inline allOf schema
	Id *string `json:"_id,omitempty" bson:"_id,omitempty"`
	// Embedded struct due to allOf(#/components/schemas/NewClinic)
	NewClinic `bson:",inline"`
}

// ClinicianPermissions defines model for ClinicianPermissions.
type ClinicianPermissions struct {
	Permissions *[]string `json:"permissions,omitempty" bson:"permissions,omitempty"`
}

// CliniciansPostid defines model for CliniciansPostid.
type CliniciansPostid struct {
	Id *string `json:"id,omitempty" bson:"id,omitempty"`
}

// ClinicsClinicians defines model for ClinicsClinicians.
type ClinicsClinicians struct {
	// Embedded fields due to inline allOf schema
	ClinicId    *string `json:"clinicId,omitempty" bson:"clinicId,omitempty"`
	ClinicianId string  `json:"clinicianId" bson:"clinicianId"`
	// Embedded struct due to allOf(#/components/schemas/ClinicianPermissions)
	ClinicianPermissions `bson:",inline"`
}

// ClinicsPatientClinician defines model for ClinicsPatientClinician.
type ClinicsPatientClinician struct {

	// Clinic Object
	Clinic     *Clinic `json:"clinic,omitempty" bson:"clinic,omitempty"`
	Clinicians *[]User `json:"clinicians,omitempty" bson:"clinicians,omitempty"`
	Patients   *[]User `json:"patients,omitempty" bson:"patients,omitempty"`
}

// ClinicsPatients defines model for ClinicsPatients.
type ClinicsPatients struct {
	// Embedded fields due to inline allOf schema
	ClinicId  *string `json:"clinicId,omitempty" bson:"clinicId,omitempty"`
	Id        *string `json:"id,omitempty" bson:"id,omitempty"`
	PatientId *string `json:"patientId,omitempty" bson:"patientId,omitempty"`
	// Embedded struct due to allOf(#/components/schemas/PatientPermissions)
	PatientPermissions `bson:",inline"`
}

// Error defines model for Error.
type Error struct {
	Code    int     `json:"code" bson:"code"`
	Message *string `json:"message,omitempty" bson:"message,omitempty"`
}

// NewClinic defines model for NewClinic.
type NewClinic struct {
	Address    *string `json:"address,omitempty" bson:"address,omitempty"`
	City       *string `json:"city,omitempty" bson:"city,omitempty"`
	ClinicSize *int    `json:"clinicSize,omitempty" bson:"clinicSize,omitempty"`
	ClinicType *string `json:"clinicType,omitempty" bson:"clinicType,omitempty"`
	Country    *string `json:"country,omitempty" bson:"country,omitempty"`
	Metadata   *struct {
		Data *string `json:"data,omitempty" bson:"data,omitempty"`
	} `json:"metadata,omitempty" bson:"metadata,omitempty"`
	Name         *string `json:"name,omitempty" bson:"name,omitempty"`
	PhoneNumbers *[]struct {
		Number *string `json:"number,omitempty" bson:"number,omitempty"`
		Type   *string `json:"type,omitempty" bson:"type,omitempty"`
	} `json:"phoneNumbers,omitempty" bson:"phoneNumbers,omitempty"`
	PostalCode *string `json:"postalCode,omitempty" bson:"postalCode,omitempty"`
	State      *string `json:"state,omitempty" bson:"state,omitempty"`
}

// PatientPermissions defines model for PatientPermissions.
type PatientPermissions struct {
	Permissions *[]string `json:"permissions,omitempty" bson:"permissions,omitempty"`
}

// PatientPostid defines model for PatientPostid.
type PatientPostid struct {
	Id *string `json:"id,omitempty" bson:"id,omitempty"`
}

// PostClinicBody defines model for PostClinicBody.
type PostClinicBody struct {
	Id *string `json:"id,omitempty" bson:"id,omitempty"`
}

// PostClinicResponse defines model for PostClinicResponse.
type PostClinicResponse struct {
	ClinicId *string `json:"clinicId,omitempty" bson:"clinicId,omitempty"`
}

// User defines model for User.
type User struct {
	Fullname string `json:"fullname" bson:"fullname"`
	Userid   string `json:"userid" bson:"userid"`
	Username string `json:"username" bson:"username"`
}

// GetClinicsParams defines parameters for GetClinics.
type GetClinicsParams struct {
	Offset      *int    `json:"offset,omitempty" bson:"offset,omitempty"`
	Limit       *int    `json:"limit,omitempty" bson:"limit,omitempty"`
	SortOrder   *string `json:"sortOrder,omitempty" bson:"sortOrder,omitempty"`
	ClinicianId *string `json:"clinicianId,omitempty" bson:"clinicianId,omitempty"`
	PatientId   *string `json:"patientId,omitempty" bson:"patientId,omitempty"`
}

// PostClinicsJSONBody defines parameters for PostClinics.
type PostClinicsJSONBody PostClinicBody

// PostClinicsParams defines parameters for PostClinics.
type PostClinicsParams struct {

	// Userid for endpoint
	XTIDEPOOLUSERID *string `json:"X-TIDEPOOL-USERID,omitempty" bson:"X-TIDEPOOL-USERID,omitempty"`
}

// GetClinicsClinicidParams defines parameters for GetClinicsClinicid.
type GetClinicsClinicidParams struct {
	Clinicians        *bool   `json:"clinicians,omitempty" bson:"clinicians,omitempty"`
	Patients          *bool   `json:"patients,omitempty" bson:"patients,omitempty"`
	IncludeClinicians *bool   `json:"includeClinicians,omitempty" bson:"includeClinicians,omitempty"`
	IncludePatients   *string `json:"includePatients,omitempty" bson:"includePatients,omitempty"`
}

// PatchClinicsClinicidJSONBody defines parameters for PatchClinicsClinicid.
type PatchClinicsClinicidJSONBody NewClinic

// GetClinicsClinicidCliniciansParams defines parameters for GetClinicsClinicidClinicians.
type GetClinicsClinicidCliniciansParams struct {
	Offset    *int    `json:"offset,omitempty" bson:"offset,omitempty"`
	Limit     *int    `json:"limit,omitempty" bson:"limit,omitempty"`
	SortOrder *string `json:"sortOrder,omitempty" bson:"sortOrder,omitempty"`
}

// PostClinicsClinicidCliniciansJSONBody defines parameters for PostClinicsClinicidClinicians.
type PostClinicsClinicidCliniciansJSONBody ClinicsClinicians

// PatchClinicsClinicidCliniciansClinicianidJSONBody defines parameters for PatchClinicsClinicidCliniciansClinicianid.
type PatchClinicsClinicidCliniciansClinicianidJSONBody ClinicianPermissions

// GetClinicsClinicidPatientsParams defines parameters for GetClinicsClinicidPatients.
type GetClinicsClinicidPatientsParams struct {
	Offset    *int    `json:"offset,omitempty" bson:"offset,omitempty"`
	Limit     *int    `json:"limit,omitempty" bson:"limit,omitempty"`
	SortOrder *string `json:"sortOrder,omitempty" bson:"sortOrder,omitempty"`
}

// PostClinicsClinicidPatientsJSONBody defines parameters for PostClinicsClinicidPatients.
type PostClinicsClinicidPatientsJSONBody ClinicsPatients

// PatchClinicsClinicidPatientsPatientidJSONBody defines parameters for PatchClinicsClinicidPatientsPatientid.
type PatchClinicsClinicidPatientsPatientidJSONBody PatientPermissions

// PostClinicsRequestBody defines body for PostClinics for application/json ContentType.
type PostClinicsJSONRequestBody PostClinicsJSONBody

// PatchClinicsClinicidRequestBody defines body for PatchClinicsClinicid for application/json ContentType.
type PatchClinicsClinicidJSONRequestBody PatchClinicsClinicidJSONBody

// PostClinicsClinicidCliniciansRequestBody defines body for PostClinicsClinicidClinicians for application/json ContentType.
type PostClinicsClinicidCliniciansJSONRequestBody PostClinicsClinicidCliniciansJSONBody

// PatchClinicsClinicidCliniciansClinicianidRequestBody defines body for PatchClinicsClinicidCliniciansClinicianid for application/json ContentType.
type PatchClinicsClinicidCliniciansClinicianidJSONRequestBody PatchClinicsClinicidCliniciansClinicianidJSONBody

// PostClinicsClinicidPatientsRequestBody defines body for PostClinicsClinicidPatients for application/json ContentType.
type PostClinicsClinicidPatientsJSONRequestBody PostClinicsClinicidPatientsJSONBody

// PatchClinicsClinicidPatientsPatientidRequestBody defines body for PatchClinicsClinicidPatientsPatientid for application/json ContentType.
type PatchClinicsClinicidPatientsPatientidJSONRequestBody PatchClinicsClinicidPatientsPatientidJSONBody

