package read

import (
	"KillReall666/schooldocumentmanagment.git/internal/model"
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type materialReader interface {
	ReadMaterialByUUID(ctx context.Context, UUID string) (*model.Publication, error)
}

type materialReadHandler struct {
	materialRead materialReader
}

func NewCreateHandler(create materialReader) *materialReadHandler {
	return &materialReadHandler{
		materialRead: create,
	}
}

func (h *materialReadHandler) Read(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "only POST method is supported", http.StatusMethodNotAllowed)
	}

	var UUIDString string
	ctx := r.Context()

	data, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("error reading body:", err)
		return
	}

	err = json.Unmarshal(data, &UUIDString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("error unmarshalling body:", err)
		return
	}

	publication, err := h.materialRead.ReadMaterialByUUID(ctx, UUIDString)
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
