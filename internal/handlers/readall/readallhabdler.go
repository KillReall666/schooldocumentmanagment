package readall

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"KillReall666/schooldocumentmanagment.git/internal/model"
)

//go:generate go run github.com/vektra/mockery/v2@v2.43.2 --name=allPublicationsReader
type allPublicationsReader interface {
	ReadAllPublicationsByUUID(ctx context.Context) ([]*model.Publication, error)
}

type allPublicationsReadHandler struct {
	allPublicationsRead allPublicationsReader
}

func NewAllPublicationsHandler(create allPublicationsReader) *allPublicationsReadHandler {
	return &allPublicationsReadHandler{
		allPublicationsRead: create,
	}
}

func (h *allPublicationsReadHandler) ReadAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "only POST method is supported", http.StatusMethodNotAllowed)
	}

	ctx := context.Background()

	publication, err := h.allPublicationsRead.ReadAllPublicationsByUUID(ctx)
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
