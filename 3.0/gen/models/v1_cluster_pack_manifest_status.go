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

// V1ClusterPackManifestStatus v1 cluster pack manifest status
//
// swagger:model v1ClusterPackManifestStatus
type V1ClusterPackManifestStatus struct {

	// condition
	Condition *V1ClusterCondition `json:"condition,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// uid
	UID string `json:"uid,omitempty"`
}

// Validate validates this v1 cluster pack manifest status
func (m *V1ClusterPackManifestStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCondition(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ClusterPackManifestStatus) validateCondition(formats strfmt.Registry) error {
	if swag.IsZero(m.Condition) { // not required
		return nil
	}

	if m.Condition != nil {
		if err := m.Condition.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("condition")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("condition")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1 cluster pack manifest status based on the context it is used
func (m *V1ClusterPackManifestStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCondition(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ClusterPackManifestStatus) contextValidateCondition(ctx context.Context, formats strfmt.Registry) error {

	if m.Condition != nil {
		if err := m.Condition.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("condition")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("condition")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1ClusterPackManifestStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ClusterPackManifestStatus) UnmarshalBinary(b []byte) error {
	var res V1ClusterPackManifestStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
