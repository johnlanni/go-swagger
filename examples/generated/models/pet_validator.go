package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/casualjim/go-swagger/errors"
	"github.com/casualjim/go-swagger/strfmt"
	"github.com/casualjim/go-swagger/validation"
)

// Validate validates this pet
func (m *Pet) Validate(formats strfmt.Registry) error {

	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePhotoUrls(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}

	return nil
}

func (m *Pet) validateName(formats strfmt.Registry) error {

	if err := validation.Required("name", "", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *Pet) validatePhotoUrls(formats strfmt.Registry) error {

	// single schema m.PhotoUrls
	for i := 0; i < len(m.PhotoUrls); i++ {

	}

	return nil
}

func (m *Pet) validateTags(formats strfmt.Registry) error {

	// single schema m.Tags
	for i := 0; i < len(m.Tags); i++ {

		// custom object Tag
		if err := m.Tags[i].Validate(formats); err != nil {
			return err
		}

	}

	return nil
}
