package create

import (
	"errors"

	"github.com/KillReall666/schooldocumentmanagment/internal/model"
)

func ValidatePublication(pub model.CreatePublication) error {
	if pub.MaterialType == "" {
		return errors.New("material_type is required")
	}
	if pub.Status == "" {
		return errors.New("status is required")
	}
	if pub.Title == "" {
		return errors.New("title is required")
	}
	if pub.Content == "" {
		return errors.New("content is required")
	}
	return nil
}
