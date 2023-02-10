// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PostParamsBodyFilesItems post params body files items
//
// swagger:model postParamsBodyFilesItems
type PostParamsBodyFilesItems struct {

	// Path to the file provided by Direktiv
	// Required: true
	File *string `json:"file"`

	// Mimetype of the file
	// Example: image/png
	// Required: true
	Mime *string `json:"mime"`

	// Name of the parameter in replicate.com API for the model
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this post params body files items
func (m *PostParamsBodyFilesItems) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFile(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostParamsBodyFilesItems) validateFile(formats strfmt.Registry) error {

	if err := validate.Required("file", "body", m.File); err != nil {
		return err
	}

	return nil
}

func (m *PostParamsBodyFilesItems) validateMime(formats strfmt.Registry) error {

	if err := validate.Required("mime", "body", m.Mime); err != nil {
		return err
	}

	return nil
}

func (m *PostParamsBodyFilesItems) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this post params body files items based on context it is used
func (m *PostParamsBodyFilesItems) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PostParamsBodyFilesItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostParamsBodyFilesItems) UnmarshalBinary(b []byte) error {
	var res PostParamsBodyFilesItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
