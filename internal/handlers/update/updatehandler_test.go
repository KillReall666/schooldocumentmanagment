package update

import (
	"net/http"
	"testing"
)

func Test_publicationUpdateHandler_Update(t *testing.T) {
	type fields struct {
		publicationUpdate publicationUpdater
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
			h := &publicationUpdateHandler{
				publicationUpdate: tt.fields.publicationUpdate,
			}
			h.Update(tt.args.w, tt.args.r)
		})
	}
}
