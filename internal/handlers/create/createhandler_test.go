package create

import (
	"KillReall666/schooldocumentmanagment.git/internal/handlers/create/mocks"
	"KillReall666/schooldocumentmanagment.git/internal/model"
	"bytes"
	"context"
	"encoding/json"
	"github.com/stretchr/testify/mock"
	"net/http"
	"strings"
	"testing"
	"time"
)

type MockResponseWriter struct {
	HeaderMap  http.Header
	StatusCode int
	Body       *bytes.Buffer
}

func (m *MockResponseWriter) Header() http.Header {
	return m.HeaderMap
}

func (m *MockResponseWriter) Write(b []byte) (int, error) {
	return m.Body.Write(b)
}

func (m *MockResponseWriter) WriteHeader(statusCode int) {
	m.StatusCode = statusCode
}

func NewMockResponseWriter() *MockResponseWriter {
	return &MockResponseWriter{
		HeaderMap:  make(http.Header),
		Body:       new(bytes.Buffer),
		StatusCode: 200, // Значение по умолчанию
	}
}

func Test_materialCreateHandler_Create(t *testing.T) {
	type fields struct {
		publicationCreater publicationCreater
	}
	type args struct {
		w           http.ResponseWriter
		r           *http.Request
		publication model.Publication
		ctx         context.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "create success",
			fields: fields{
				publicationCreater: mocks.NewPublicationCreater(t),
			},
			args: args{
				w: NewMockResponseWriter(),
				r: func() *http.Request {
					publication := model.CreatePublication{
						MaterialType: "Article",
						Status:       "Published",
						Title:        "Введение в программирование на Python",
						Content:      "Python — это высокоуровневый язык программирования, который позволяет быстро и эффективно разрабатывать приложения.",
						CreatedAt:    time.Date(2023, 10, 5, 12, 30, 0, 0, time.UTC),
						UpdatedAt:    time.Date(2023, 10, 5, 12, 30, 0, 0, time.UTC),
					}
					body, _ := json.Marshal(publication)
					req, _ := http.NewRequest("POST", "/api/articles", strings.NewReader(string(body)))
					req.Header.Set("Content-Type", "application/json")
					return req
				}(),
				publication: model.Publication{
					MaterialType: "Article",
					Status:       "Published",
					Title:        "Введение в программирование на Python",
					Content:      "Python — это высокоуровневый язык программирования, который позволяет быстро и эффективно разрабатывать приложения.",
					CreatedAt:    time.Date(2023, 10, 5, 12, 30, 0, 0, time.UTC),
					UpdatedAt:    time.Date(2023, 10, 5, 12, 30, 0, 0, time.UTC),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			createPublication := mocks.NewPublicationCreater(t)
			ctx := context.Background()

			createPublication.
				On("CreatePublication", ctx, mock.Anything, tt.args.publication.MaterialType, tt.args.publication.Status, tt.args.publication.Title, tt.args.publication.Content, tt.args.publication.CreatedAt, tt.args.publication.UpdatedAt).
				Return(nil)

			h := &publicationCreateHandler{
				createPublication,
			}

			h.Create(tt.args.w, tt.args.r)
		})
	}
}
