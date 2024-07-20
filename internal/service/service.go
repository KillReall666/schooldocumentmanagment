package service

import (
	"KillReall666/schooldocumentmanagment.git/internal/config"
	"KillReall666/schooldocumentmanagment.git/internal/model"
	"KillReall666/schooldocumentmanagment.git/internal/storage/postgres"
	"context"
	"github.com/google/uuid"
	"time"
)

type service struct {
	cfg *config.Config
	//log *logger.Logger
	db *postgres.Database
}

type Service interface {
}

func New(cfg *config.Config, db *postgres.Database) *service { //log *logger.Logger,
	return &service{
		cfg: cfg,
		//	log: log,
		db: db,
	}
}

func (s *service) CreatePublication(ctx context.Context, ID uuid.UUID, MaterialType string, Status string, Title string, Content string, CreatedAt time.Time, UpdatedAt time.Time) error {
	err := s.db.CreatePublication(ctx, ID, MaterialType, Status, Title, Content, CreatedAt, UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) ReadMaterialByUUID(ctx context.Context, UUID string) (*model.Publication, error) {
	data, err := s.db.ReadMaterialByUUID(ctx, UUID)
	if err != nil {
		return nil, err
	}
	return data, nil

}

func (s *service) UpdateMaterialByUUID(ctx context.Context, publications model.Publication) error {
	err := s.db.UpdateMaterialByUUID(ctx, publications)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) ReadAllMaterialByUUID(ctx context.Context) ([]*model.Publication, error) {
	data, err := s.db.ReadAllMaterialByUUID(ctx)
	if err != nil {
		return nil, err
	}
	return data, nil
}
