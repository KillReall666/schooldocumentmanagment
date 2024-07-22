package create

import (
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"io"
	"net/http"

	"github.com/KillReall666/schooldocumentmanagment/internal/model"
)

//go:generate go run github.com/vektra/mockery/v2@v2.43.2 --name=publicationCreater
type publicationCreater interface {
	CreatePublication(ctx context.Context, ID uuid.UUID, material model.CreatePublication) error
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
		return
	}

	var publication model.CreatePublication
	ctx := context.Background()

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

	if err = ValidatePublication(publication); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newUUID := uuid.New()

	err = h.publicationCreate.CreatePublication(ctx, newUUID, publication)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(newUUID.String()))
}
