package service

import (
	"context"
	"github.com/google/uuid"

	"github.com/KillReall666/schooldocumentmanagment/internal/config"
	"github.com/KillReall666/schooldocumentmanagment/internal/model"
	"github.com/KillReall666/schooldocumentmanagment/internal/storage/postgres"
)

type service struct {
	cfg *config.Config
	db  postgres.PublicationsRepository
}

func New(cfg *config.Config, db postgres.PublicationsRepository) *service {
	return &service{
		cfg: cfg,
		db:  db,
	}
}

func (s *service) CreatePublication(ctx context.Context, ID uuid.UUID, material model.CreatePublication) error {
	err := s.db.CreatePublication(ctx, ID, material)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) ReadPublicationByUUID(ctx context.Context, UUID string) (*model.Publication, error) {
	data, err := s.db.ReadPublicationByUUID(ctx, UUID)
	if err != nil {
		return nil, err
	}
	return data, nil

}

func (s *service) UpdatePublicationByUUID(ctx context.Context, publications model.Publication) error {
	err := s.db.UpdatePublicationByUUID(ctx, publications)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) ReadAllPublicationsByUUID(ctx context.Context) ([]*model.Publication, error) {
	data, err := s.db.ReadAllPublicationsByUUID(ctx)
	if err != nil {
		return nil, err
	}
	return data, nil
}
