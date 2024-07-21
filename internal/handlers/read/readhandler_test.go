package read

import (
	"net/http"
	"testing"
)

func Test_publicationReadHandler_Read(t *testing.T) {
	type fields struct {
		publicationRead publicationReader
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
			h := &publicationReadHandler{
				publicationRead: tt.fields.publicationRead,
			}
			h.Read(tt.args.w, tt.args.r)
		})
	}
}
