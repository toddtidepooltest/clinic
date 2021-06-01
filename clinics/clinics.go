package clinics

import (
	"context"
	"fmt"
	"github.com/tidepool-org/clinic/errors"
	"github.com/tidepool-org/clinic/store"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

var ErrNotFound = fmt.Errorf("clinic %w", errors.NotFound)
var ErrDuplicateShareCode = fmt.Errorf("%w share code", errors.Duplicate)
var ErrAdminRequired = fmt.Errorf("%w: the clinic must have at least one admin", errors.ConstraintViolation)

type Service interface {
	Get(ctx context.Context, id string) (*Clinic, error)
	List(ctx context.Context, filter *Filter, pagination store.Pagination) ([]*Clinic, error)
	Create(ctx context.Context, clinic *Clinic) (*Clinic, error)
	Update(ctx context.Context, id string, clinic *Clinic) (*Clinic, error)
	UpsertAdmin(ctx context.Context, clinicId, clinicianId string) error
	RemoveAdmin(ctx context.Context, clinicId, clinicianId string) error
}

type Filter struct {
	Ids        []string
	Email      *string
	ShareCodes []string
}

type Clinic struct {
	Id                 *primitive.ObjectID `bson:"_id,omitempty"`
	Address            *string             `bson:"address,omitempty"`
	City               *string             `bson:"city,omitempty"`
	ClinicType         *string             `bson:"clinicType,omitempty"`
	Country            *string             `bson:"country,omitempty"`
	Name               *string             `bson:"name,omitempty"`
	PhoneNumbers       *[]PhoneNumber      `bson:"phoneNumbers,omitempty"`
	PostalCode         *string             `bson:"postalCode,omitempty"`
	State              *string             `bson:"state,omitempty"`
	CanonicalShareCode *string             `bson:"canonicalShareCode,omitempty"`
	ShareCodes         *[]string           `bson:"shareCodes,omitempty"`
	Admins             *[]string           `bson:"admins,omitempty"`
	CreatedAt          time.Time           `bson:"createdAt"`
	UpdatedAt          time.Time           `bson:"updatedAt"`
}

type PhoneNumber struct {
	Type   *string `bson:"type,omitempty"`
	Number string  `bson:"number,omitempty"`
}
