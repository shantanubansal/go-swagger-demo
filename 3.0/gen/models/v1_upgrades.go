// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1Upgrades Upgrades represent the reason of the last upgrade that took place
//
// swagger:model v1Upgrades
type V1Upgrades struct {

	// reason
	Reason []string `json:"reason"`

	// timestamp
	// Format: date-time
	Timestamp V1Time `json:"timestamp,omitempty"`
}

// Validate validates this v1 upgrades
func (m *V1Upgrades) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1Upgrades) validateTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.Timestamp) { // not required
		return nil
	}

	if err := m.Timestamp.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("timestamp")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("timestamp")
		}
		return err
	}

	return nil
}

// ContextValidate validate this v1 upgrades based on the context it is used
func (m *V1Upgrades) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTimestamp(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1Upgrades) contextValidateTimestamp(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Timestamp.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("timestamp")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("timestamp")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1Upgrades) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1Upgrades) UnmarshalBinary(b []byte) error {
	var res V1Upgrades
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
