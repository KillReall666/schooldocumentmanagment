package readall

import (
	"KillReall666/schooldocumentmanagment.git/internal/model"
	"context"
	"encoding/json"
	"log"
	"net/http"
)

type materialAllReader interface {
	ReadAllMaterialByUUID(ctx context.Context) ([]*model.Publication, error)
}

type materialAllReadHandler struct {
	materialAllRead materialAllReader
}

func NewCreateHandler(create materialAllReader) *materialAllReadHandler {
	return &materialAllReadHandler{
		materialAllRead: create,
	}
}

func (h *materialAllReadHandler) ReadAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "only POST method is supported", http.StatusMethodNotAllowed)
	}

	ctx := r.Context()

	publication, err := h.materialAllRead.ReadAllMaterialByUUID(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("error reading material:", err)
		return
	}

	marshalData, err := json.Marshal(publication)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("error marshaling material:", err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	w.Write(marshalData)
}
