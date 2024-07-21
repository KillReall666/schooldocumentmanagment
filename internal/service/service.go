package service

import (
	"context"
	"time"

	"github.com/google/uuid"

	"KillReall666/schooldocumentmanagment.git/internal/config"
	"KillReall666/schooldocumentmanagment.git/internal/model"
	"KillReall666/schooldocumentmanagment.git/internal/storage/postgres"
)

type service struct {
	cfg *config.Config
	db  *postgres.Database
}

type Service interface {
}

func New(cfg *config.Config, db *postgres.Database) *service {
	return &service{
		cfg: cfg,
		db:  db,
	}
}

func (s *service) CreatePublication(ctx context.Context, ID uuid.UUID, MaterialType string, Status string, Title string, Content string, CreatedAt time.Time, UpdatedAt time.Time) error {
	err := s.db.CreatePublication(ctx, ID, MaterialType, Status, Title, Content, CreatedAt, UpdatedAt)
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
