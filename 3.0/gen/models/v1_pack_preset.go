// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1PackPreset PackPreset defines the preset pack values
//
// swagger:model v1PackPreset
type V1PackPreset struct {

	// add
	Add string `json:"add"`

	// display name
	DisplayName string `json:"displayName"`

	// group
	Group string `json:"group"`

	// name
	Name string `json:"name"`

	// remove
	Remove []string `json:"remove"`
}

// Validate validates this v1 pack preset
func (m *V1PackPreset) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 pack preset based on context it is used
func (m *V1PackPreset) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1PackPreset) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1PackPreset) UnmarshalBinary(b []byte) error {
	var res V1PackPreset
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
