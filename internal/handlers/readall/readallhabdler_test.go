package readall

import (
	"net/http"
	"testing"
)

func Test_allPublicationsReadHandler_ReadAll(t *testing.T) {
	type fields struct {
		allPublicationsRead allPublicationsReader
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &allPublicationsReadHandler{
				allPublicationsRead: tt.fields.allPublicationsRead,
			}
			h.ReadAll(tt.args.w, tt.args.r)
		})
	}
}
