package api

import (
	"github.com/labstack/echo/v4"
	"github.com/tidepool-org/clinic/cliniccreator"
	"github.com/tidepool-org/clinic/clinicians"
	"github.com/tidepool-org/clinic/clinics"
	"github.com/tidepool-org/clinic/patients"
	"github.com/tidepool-org/clinic/store"
	"go.uber.org/fx"
)

type Handler struct {
	clinics           clinics.Service
	clinicsCreator    cliniccreator.Creator
	clinicians        clinicians.Service
	cliniciansUpdater clinicians.Service
	patients          patients.Service
	users             patients.UserService
}

var _ ServerInterface = &Handler{}

type Params struct {
	fx.In

	Clinics           clinics.Service
	ClinicsCreator    cliniccreator.Creator
	Clinicians        clinicians.Service
	CliniciansUpdater clinicians.Service
	Patients          patients.Service
	Users             patients.UserService
}

func NewHandler(p Params) *Handler {
	return &Handler{
		clinics:           p.Clinics,
		clinicsCreator:    p.ClinicsCreator,
		clinicians:        p.Clinicians,
		cliniciansUpdater: p.CliniciansUpdater,
		patients:          p.Patients,
		users:             p.Users,
	}
}

func pagination(offset *Offset, limit *Limit) store.Pagination {
	page := store.DefaultPagination()
	if offset != nil {
		page.Offset = int(*offset)
	}
	if limit != nil {
		page.Limit = int(*limit)
	}
	return page
}
