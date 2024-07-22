package update

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/KillReall666/schooldocumentmanagment/internal/model"
)

//go:generate go run github.com/vektra/mockery/v2@v2.43.2 --name=publicationUpdater
type publicationUpdater interface {
	UpdatePublicationByUUID(ctx context.Context, publications model.Publication) error
}

type publicationUpdateHandler struct {
	publicationUpdate publicationUpdater
}

func NewUpdateHandler(update publicationUpdater) *publicationUpdateHandler {
	return &publicationUpdateHandler{
		publicationUpdate: update,
	}
}

func (h *publicationUpdateHandler) Update(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "only POST method is supported", http.StatusMethodNotAllowed)
	}

	var publication model.Publication
	ctx := context.Background()

	data, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("error reading body:", err)
		return
	}

	err = json.Unmarshal(data, &publication)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("error unmarshalling body:", err)
		return
	}

	err = h.publicationUpdate.UpdatePublicationByUUID(ctx, publication)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("error reading material:", err)
		return
	}

	marshalData, err := json.Marshal(publication)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("error marshalling material:", err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	w.Write(marshalData)
}
