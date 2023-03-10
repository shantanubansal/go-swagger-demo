// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1EventSource Event source info
//
// swagger:model v1EventSource
type V1EventSource struct {

	// component
	Component string `json:"component,omitempty"`

	// host
	Host string `json:"host,omitempty"`
}

// Validate validates this v1 event source
func (m *V1EventSource) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1EventSource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1EventSource) UnmarshalBinary(b []byte) error {
	var res V1EventSource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
