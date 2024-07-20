package update

import (
	"KillReall666/schooldocumentmanagment.git/internal/model"
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type materialUpdater interface {
	UpdateMaterialByUUID(ctx context.Context, publications model.Publication) error
}

type materialUpdateHandler struct {
	materialUpdate materialUpdater
}

func NewUpdateHandler(update materialUpdater) *materialUpdateHandler {
	return &materialUpdateHandler{
		materialUpdate: update,
	}
}

func (h *materialUpdateHandler) Update(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "only POST method is supported", http.StatusMethodNotAllowed)
	}

	var publication model.Publication
	ctx := r.Context()

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

	err = h.materialUpdate.UpdateMaterialByUUID(ctx, publication)
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
