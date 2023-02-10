// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PostParamsBodyFiles post params body files
//
// swagger:model postParamsBodyFiles
type PostParamsBodyFiles struct {

	// file
	File string `json:"file,omitempty"`

	// mime
	Mime string `json:"mime,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this post params body files
func (m *PostParamsBodyFiles) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this post params body files based on context it is used
func (m *PostParamsBodyFiles) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PostParamsBodyFiles) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostParamsBodyFiles) UnmarshalBinary(b []byte) error {
	var res PostParamsBodyFiles
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
