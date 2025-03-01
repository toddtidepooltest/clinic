package patients

import (
	"context"
	"errors"
	"github.com/tidepool-org/clinic/store"
	"go.uber.org/zap"
)

type service struct {
	repo             Repository
	custodialService CustodialService
	logger           *zap.SugaredLogger
}

var _ Service = &service{}

func NewService(repo Repository, custodialService CustodialService, logger *zap.SugaredLogger) (Service, error) {
	return &service{
		repo:             repo,
		custodialService: custodialService,
		logger:           logger,
	}, nil
}

func (s *service) Get(ctx context.Context, clinicId string, userId string) (*Patient, error) {
	return s.repo.Get(ctx, clinicId, userId)
}

func (s *service) List(ctx context.Context, filter *Filter, pagination store.Pagination, sort *store.Sort) (*ListResult, error) {
	return s.repo.List(ctx, filter, pagination, sort)
}

func (s *service) Create(ctx context.Context, patient Patient) (*Patient, error) {
	// Only create new accounts if the custodial user doesn't exist already (i.e. we are not migrating it)
	if patient.IsCustodial() && patient.UserId == nil {
		s.logger.Infow("creating custodial account", "clinicId", patient.ClinicId.Hex())
		return s.custodialService.CreateAccount(ctx, patient)
	}
	if patient.UserId == nil {
		return nil, errors.New("user id is missing")
	}
	s.logger.Infow("creating patient in clinic", "userId", patient.UserId, "clinicId", patient.ClinicId.Hex())
	return s.repo.Create(ctx, patient)
}

func (s *service) Update(ctx context.Context, clinicId string, userId string, patient Patient) (*Patient, error) {
	existing, err := s.Get(ctx, clinicId, userId)
	if err != nil {
		return nil, err
	}

	if existing.IsCustodial() {
		s.logger.Infow("updating custodial account", "userId", existing.UserId, "clinicId", clinicId)
		patient.ClinicId = existing.ClinicId
		patient.UserId = existing.UserId
		return s.custodialService.UpdateAccount(ctx, patient)
	}

	s.logger.Infow("updating patient", "userId", existing.UserId, "clinicId", clinicId)
	return s.repo.Update(ctx, clinicId, userId, patient)
}

func (s *service) Remove(ctx context.Context, clinicId string, userId string) error {
	s.logger.Infow("deleting patient from clinic", "userId", userId, "clinicId", clinicId)
	return s.repo.Remove(ctx, clinicId, userId)
}

func (s *service) UpdatePermissions(ctx context.Context, clinicId, userId string, permissions *Permissions) (*Patient, error) {
	if permissions != nil && permissions.Custodian != nil {
		// Custodian permission cannot be set after patients claimed their accounts
		permissions.Custodian = nil
	}
	if permissions == nil || permissions.Empty() {
		s.logger.Infow(
			"deleting patient from clinic because the patient revoked all permissions",
			"userId", userId, "clinicId", clinicId,
		)
		return nil, s.Remove(ctx, clinicId, userId)
	}
	return s.repo.UpdatePermissions(ctx, clinicId, userId, permissions)
}

func (s *service) DeletePermission(ctx context.Context, clinicId, userId, permission string) (*Patient, error) {
	patient, err := s.repo.DeletePermission(ctx, clinicId, userId, permission)
	if err != nil {
		return nil, err
	}
	if shouldRemovePatientFromClinic(patient) {
		s.logger.Infow(
			"deleting patient from clinic because the patient revoked all permissions",
			"userId", userId, "clinicId", clinicId,
		)
		if err := s.Remove(ctx, clinicId, userId); err != nil {
			// the patient was removed by concurrent request which is not a problem,
			// because it had to be removed as a result of the current operation
			if err == ErrNotFound {
				return nil, nil
			}
			return nil, err
		}
		return nil, nil
	}
	return patient, err
}

func (s *service) DeleteFromAllClinics(ctx context.Context, userId string) error {
	s.logger.Infow("deleting patients from all clinics", "userId", userId)
	return s.repo.DeleteFromAllClinics(ctx, userId)
}

func (s *service) DeleteNonCustodialPatientsOfClinic(ctx context.Context, clinicId string) error {
	s.logger.Infow("deleting all non-custodial patient of clinic", "clinicId", clinicId)
	return s.repo.DeleteNonCustodialPatientsOfClinic(ctx, clinicId)
}

func shouldRemovePatientFromClinic(patient *Patient) bool {
	if patient != nil {
		return patient.Permissions == nil || patient.Permissions.Empty()
	}
	return false
}
