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

// L7 Auth token response
//
// swagger:model l7
type L7 struct {

	// arr
	Arr *ArrSpec `json:"arr,omitempty"`

	// l8
	L8 *L8 `json:"l8,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this l7
func (m *L7) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateArr(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateL8(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *L7) validateArr(formats strfmt.Registry) error {
	if swag.IsZero(m.Arr) { // not required
		return nil
	}

	if m.Arr != nil {
		if err := m.Arr.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("arr")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("arr")
			}
			return err
		}
	}

	return nil
}

func (m *L7) validateL8(formats strfmt.Registry) error {
	if swag.IsZero(m.L8) { // not required
		return nil
	}

	if m.L8 != nil {
		if err := m.L8.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("l8")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("l8")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this l7 based on the context it is used
func (m *L7) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateArr(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateL8(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *L7) contextValidateArr(ctx context.Context, formats strfmt.Registry) error {

	if m.Arr != nil {
		if err := m.Arr.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("arr")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("arr")
			}
			return err
		}
	}

	return nil
}

func (m *L7) contextValidateL8(ctx context.Context, formats strfmt.Registry) error {

	if m.L8 != nil {
		if err := m.L8.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("l8")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("l8")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *L7) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *L7) UnmarshalBinary(b []byte) error {
	var res L7
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
