package create

import (
	"testing"

	"github.com/KillReall666/schooldocumentmanagment/internal/model"
)

func TestValidatePublication(t *testing.T) {
	type args struct {
		pub model.CreatePublication
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "valid publication",
			args: args{
				pub: model.CreatePublication{
					MaterialType: "Article",
					Status:       "Published",
					Title:        "Test Title",
					Content:      "Test Content",
				},
			},
			wantErr: false,
		},
		{
			name: "missing material_type",
			args: args{
				pub: model.CreatePublication{
					Status:  "Published",
					Title:   "Test Title",
					Content: "Test Content",
				},
			},
			wantErr: true,
		},
		{
			name: "missing status",
			args: args{
				pub: model.CreatePublication{
					MaterialType: "Article",
					Title:        "Test Title",
					Content:      "Test Content",
				},
			},
			wantErr: true,
		},
		{
			name: "missing title",
			args: args{
				pub: model.CreatePublication{
					MaterialType: "Article",
					Status:       "Published",
					Content:      "Test Content",
				},
			},
			wantErr: true,
		},
		{
			name: "missing content",
			args: args{
				pub: model.CreatePublication{
					MaterialType: "Article",
					Status:       "Published",
					Title:        "Test Title",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidatePublication(tt.args.pub); (err != nil) != tt.wantErr {
				t.Errorf("ValidatePublication() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
