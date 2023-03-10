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

// Status Auth token response
//
// swagger:model status
type Status struct {

	// creation time
	// Format: date-time
	CreationTime V1Time `json:"creationTime,omitempty"`

	// last login
	// Format: date-time
	LastLogin V1Time `json:"lastLogin,omitempty"`

	// modification time
	// Format: date-time
	ModificationTime V1Time `json:"modificationTime,omitempty"`
}

// Validate validates this status
func (m *Status) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreationTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastLogin(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModificationTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Status) validateCreationTime(formats strfmt.Registry) error {
	if swag.IsZero(m.CreationTime) { // not required
		return nil
	}

	if err := m.CreationTime.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("creationTime")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("creationTime")
		}
		return err
	}

	return nil
}

func (m *Status) validateLastLogin(formats strfmt.Registry) error {
	if swag.IsZero(m.LastLogin) { // not required
		return nil
	}

	if err := m.LastLogin.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("lastLogin")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("lastLogin")
		}
		return err
	}

	return nil
}

func (m *Status) validateModificationTime(formats strfmt.Registry) error {
	if swag.IsZero(m.ModificationTime) { // not required
		return nil
	}

	if err := m.ModificationTime.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("modificationTime")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("modificationTime")
		}
		return err
	}

	return nil
}

// ContextValidate validate this status based on the context it is used
func (m *Status) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCreationTime(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLastLogin(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateModificationTime(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Status) contextValidateCreationTime(ctx context.Context, formats strfmt.Registry) error {

	if err := m.CreationTime.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("creationTime")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("creationTime")
		}
		return err
	}

	return nil
}

func (m *Status) contextValidateLastLogin(ctx context.Context, formats strfmt.Registry) error {

	if err := m.LastLogin.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("lastLogin")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("lastLogin")
		}
		return err
	}

	return nil
}

func (m *Status) contextValidateModificationTime(ctx context.Context, formats strfmt.Registry) error {

	if err := m.ModificationTime.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("modificationTime")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("modificationTime")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Status) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Status) UnmarshalBinary(b []byte) error {
	var res Status
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
