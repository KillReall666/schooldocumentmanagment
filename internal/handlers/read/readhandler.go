package read

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/KillReall666/schooldocumentmanagment/internal/model"
)

//go:generate go run github.com/vektra/mockery/v2@v2.43.2 --name=publicationReader
type publicationReader interface {
	ReadPublicationByUUID(ctx context.Context, UUID string) (*model.Publication, error)
}

type publicationReadHandler struct {
	publicationRead publicationReader
}

func NewReadHandler(create publicationReader) *publicationReadHandler {
	return &publicationReadHandler{
		publicationRead: create,
	}
}

func (h *publicationReadHandler) Read(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "only POST method is supported", http.StatusMethodNotAllowed)
	}

	var UUIDString string
	ctx := context.Background()

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

	publication, err := h.publicationRead.ReadPublicationByUUID(ctx, UUIDString)
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
