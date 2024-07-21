package create

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"testing"
	"time"
	
	"github.com/stretchr/testify/mock"

	"KillReall666/schooldocumentmanagment.git/internal/handlers/create/mocks"
	"KillReall666/schooldocumentmanagment.git/internal/model"
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
		StatusCode: 200,
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
		name           string
		fields         fields
		args           args
		wantStatusCode int
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
						MaterialType: "test1",
						Status:       "test1",
						Title:        "test1",
						Content:      "test1",
						CreatedAt:    time.Date(2023, 10, 5, 12, 30, 0, 0, time.UTC),
						UpdatedAt:    time.Date(2023, 10, 5, 12, 30, 0, 0, time.UTC),
					}
					body, _ := json.Marshal(publication)
					req, _ := http.NewRequest("POST", "/api/articles", strings.NewReader(string(body)))
					req.Header.Set("Content-Type", "application/json")
					return req
				}(),
				publication: model.Publication{
					MaterialType: "test1",
					Status:       "test1",
					Title:        "test1",
					Content:      "test1",
					CreatedAt:    time.Date(2023, 10, 5, 12, 30, 0, 0, time.UTC),
					UpdatedAt:    time.Date(2023, 10, 5, 12, 30, 0, 0, time.UTC),
				},
			},
			wantStatusCode: http.StatusCreated,
		},
		{
			name: "create failure - invalid method",
			fields: fields{
				publicationCreater: mocks.NewPublicationCreater(t),
			},
			args: args{
				w: NewMockResponseWriter(),
				r: func() *http.Request {
					publication := model.CreatePublication{
						MaterialType: "test1",
						Status:       "test1",
						Title:        "test1",
						Content:      "test1",
						CreatedAt:    time.Date(2023, 10, 5, 12, 30, 0, 0, time.UTC),
						UpdatedAt:    time.Date(2023, 10, 5, 12, 30, 0, 0, time.UTC),
					}
					body, _ := json.Marshal(publication)
					req, _ := http.NewRequest("GET", "/api/articles", strings.NewReader(string(body)))
					req.Header.Set("Content-Type", "application/json")
					return req
				}(),
				publication: model.Publication{
					MaterialType: "test1",
					Status:       "test1",
					Title:        "test1",
					Content:      "test1",
					CreatedAt:    time.Date(2023, 10, 5, 12, 30, 0, 0, time.UTC),
					UpdatedAt:    time.Date(2023, 10, 5, 12, 30, 0, 0, time.UTC),
				},
			},
			wantStatusCode: http.StatusMethodNotAllowed,
		},
		{
			name: "create failure - invalid data",
			fields: fields{
				publicationCreater: mocks.NewPublicationCreater(t),
			},
			args: args{
				w: NewMockResponseWriter(),
				r: func() *http.Request {
					publication := model.CreatePublication{
						MaterialType: "",
						Status:       "test1",
						Title:        "test1",
						Content:      "test1",
						CreatedAt:    time.Date(2023, 10, 5, 12, 30, 0, 0, time.UTC),
						UpdatedAt:    time.Date(2023, 10, 5, 12, 30, 0, 0, time.UTC),
					}
					body, _ := json.Marshal(publication)
					req, _ := http.NewRequest("POST", "/api/articles", strings.NewReader(string(body)))
					req.Header.Set("Content-Type", "application/json")
					return req
				}(),
				publication: model.Publication{
					MaterialType: "",
					Status:       "test1",
					Title:        "test1",
					Content:      "test1",
					CreatedAt:    time.Date(2023, 10, 5, 12, 30, 0, 0, time.UTC),
					UpdatedAt:    time.Date(2023, 10, 5, 12, 30, 0, 0, time.UTC),
				},
			},
			wantStatusCode: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			createPublication := mocks.NewPublicationCreater(t)
			ctx := context.Background()

			if tt.wantStatusCode == http.StatusCreated {
				createPublication.
					On("CreatePublication", ctx, mock.Anything, tt.args.publication.MaterialType, tt.args.publication.Status, tt.args.publication.Title, tt.args.publication.Content, tt.args.publication.CreatedAt, tt.args.publication.UpdatedAt).
					Once().
					Return(nil)
			}

			h := &publicationCreateHandler{
				createPublication,
			}

			h.Create(tt.args.w, tt.args.r)

			if gotStatusCode := tt.args.w.(*MockResponseWriter).StatusCode; gotStatusCode != tt.wantStatusCode {
				t.Errorf("Create() = %v, want %v", gotStatusCode, tt.wantStatusCode)
			}
		})
	}
}
