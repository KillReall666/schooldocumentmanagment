package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/KillReall666/schooldocumentmanagment/internal/model"
)

type PublicationsRepository interface {
	CreatePublication(ctx context.Context, ID uuid.UUID, material model.CreatePublication) error
	ReadPublicationByUUID(ctx context.Context, UUID string) (*model.Publication, error)
	ReadAllPublicationsByUUID(ctx context.Context) ([]*model.Publication, error)
	UpdatePublicationByUUID(ctx context.Context, publications model.Publication) error
}

type Database struct {
	db *pgxpool.Pool
}

const createPublicationTableQuery = `
      CREATE TABLE IF NOT EXISTS publications (
	id UUID PRIMARY KEY,
    material_type VARCHAR NOT NULL,
    status VARCHAR NOT NULL,
    title VARCHAR NOT NULL,
    content TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);`

const published = "Published"

func New(connString string) (*Database, error) {
	conn, err := pgxpool.New(context.Background(), connString)
	if err != nil {
		return nil, fmt.Errorf("error connecting to db: %v", err)
	}

	_, err = conn.Exec(context.Background(), createPublicationTableQuery)
	if err != nil {
		return nil, fmt.Errorf("error creating user table: %v", err)
	}

	return &Database{db: conn}, nil
}

//TODO: по хорошему вынести коннект с базой в отдельный пакет, и сделать мини-репозетории для работы с пользователем, публикациями итд.

func (d *Database) CreatePublication(ctx context.Context, ID uuid.UUID, material model.CreatePublication) error {
	createQuery := `INSERT INTO publications (id, material_type, status, title, content, created_at, updated_at) 
					VALUES ($1, $2, $3, $4, $5, $6, $7);`
	createdAt := time.Now().UTC()
	moscowTime, err := time.LoadLocation("Europe/Moscow")

	_, err = d.db.Exec(ctx, createQuery, ID, material.MaterialType, material.Status, material.Title, material.Content, createdAt.In(moscowTime), material.UpdatedAt)
	if err != nil {
		return fmt.Errorf("error creating publication: %v", err)
	}
	return nil
}

func (d *Database) ReadPublicationByUUID(ctx context.Context, UUID string) (*model.Publication, error) {
	var data model.Publication
	readQuery := `SELECT id AS uuid, material_type, status, title, content, created_at, updated_at FROM publications WHERE id = $1`

	err := d.db.QueryRow(ctx, readQuery, UUID).Scan(&data.ID, &data.MaterialType, &data.Status, &data.Title, &data.Content, &data.CreatedAt, &data.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (d *Database) ReadAllPublicationsByUUID(ctx context.Context) ([]*model.Publication, error) {
	var data []*model.Publication

	readAllQuery := `SELECT id, material_type, title, created_at, updated_at from publications`

	rows, err := d.db.Query(ctx, readAllQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var publication model.Publication
		if err = rows.Scan(&publication.ID, &publication.MaterialType, &publication.Title, &publication.CreatedAt, &publication.UpdatedAt); err != nil {
			return nil, err
		}
		data = append(data, &publication)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return data, nil

}

func (d *Database) UpdatePublicationByUUID(ctx context.Context, publications model.Publication) error {
	tx, err := d.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return fmt.Errorf("error starting transaction: %v", err)
	}
	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		}
	}()

	var status string

	err = d.db.QueryRow(ctx, `SELECT status FROM publications WHERE id = $1`, publications.ID).Scan(&status)
	if err != nil {
		return fmt.Errorf("error updating publication: %v", err)
	}

	if status != published {
		return fmt.Errorf("publication in the archive: %s", status)
	}

	createdAt := time.Now().UTC()
	moscowTime, err := time.LoadLocation("Europe/Moscow")

	_, err = d.db.Exec(context.Background(),
		`UPDATE publications 
        	 SET status = $1, title = $2, content = $3, updated_at = $4 
         	 WHERE id = $5;`, publications.Status, publications.Title, publications.Content, createdAt.In(moscowTime), publications.ID)

	if err != nil {
		return fmt.Errorf("error updating publication: %v", err)
	}

	if err = tx.Commit(ctx); err != nil {
		return fmt.Errorf("error committing transaction: %v", err)
	}

	return nil

}
