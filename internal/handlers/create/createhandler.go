package create

import (
	"KillReall666/schooldocumentmanagment.git/internal/model"
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"io"
	"net/http"
	"time"
)

//go:generate go run github.com/vektra/mockery/v2@v2.43.2 --name=materialCreater
type publicationCreater interface {
	CreatePublication(ctx context.Context, ID uuid.UUID, MaterialType string, Status string, Title string, Content string, CreatedAt time.Time, UpdatedAt time.Time) error
}

type publicationCreateHandler struct {
	publicationCreate publicationCreater
}

func NewCreateHandler(create publicationCreater) *publicationCreateHandler {
	return &publicationCreateHandler{
		publicationCreate: create,
	}

}

func (h *publicationCreateHandler) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is supported", http.StatusMethodNotAllowed)
	}

	var publication model.CreatePublication
	ctx := r.Context()
	unmarshalData, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(unmarshalData, &publication)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	newUUID := uuid.New()
	fmt.Println(publication.MaterialType, publication.Status)
	err = h.publicationCreate.CreatePublication(ctx, newUUID, publication.MaterialType, publication.Status, publication.Title, publication.Content, publication.CreatedAt, publication.UpdatedAt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(newUUID.String()))
}
