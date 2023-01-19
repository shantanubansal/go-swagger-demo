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

// Metadata Auth token response
//
// swagger:model metadata
type Metadata struct {

	// age
	Age string `json:"age,omitempty"`

	// dob
	// Format: date-time
	Dob V1Time `json:"dob,omitempty"`

	// l1
	L1 *L1 `json:"l1,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this metadata
func (m *Metadata) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDob(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateL1(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Metadata) validateDob(formats strfmt.Registry) error {
	if swag.IsZero(m.Dob) { // not required
		return nil
	}

	if err := m.Dob.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("dob")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("dob")
		}
		return err
	}

	return nil
}

func (m *Metadata) validateL1(formats strfmt.Registry) error {
	if swag.IsZero(m.L1) { // not required
		return nil
	}

	if m.L1 != nil {
		if err := m.L1.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("l1")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("l1")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this metadata based on the context it is used
func (m *Metadata) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDob(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateL1(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Metadata) contextValidateDob(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Dob.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("dob")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("dob")
		}
		return err
	}

	return nil
}

func (m *Metadata) contextValidateL1(ctx context.Context, formats strfmt.Registry) error {

	if m.L1 != nil {
		if err := m.L1.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("l1")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("l1")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Metadata) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Metadata) UnmarshalBinary(b []byte) error {
	var res Metadata
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}