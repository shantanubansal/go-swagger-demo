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

// V1ComputeMetrics Compute metrics
//
// swagger:model v1ComputeMetrics
type V1ComputeMetrics struct {

	// last updated time
	// Format: date-time
	LastUpdatedTime V1Time `json:"lastUpdatedTime,omitempty"`

	// limit
	Limit float64 `json:"limit"`

	// request
	Request float64 `json:"request"`

	// total
	Total float64 `json:"total"`

	// unit
	Unit string `json:"unit,omitempty"`

	// usage
	Usage float64 `json:"usage"`
}

// Validate validates this v1 compute metrics
func (m *V1ComputeMetrics) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLastUpdatedTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ComputeMetrics) validateLastUpdatedTime(formats strfmt.Registry) error {
	if swag.IsZero(m.LastUpdatedTime) { // not required
		return nil
	}

	if err := m.LastUpdatedTime.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("lastUpdatedTime")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("lastUpdatedTime")
		}
		return err
	}

	return nil
}

// ContextValidate validate this v1 compute metrics based on the context it is used
func (m *V1ComputeMetrics) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLastUpdatedTime(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ComputeMetrics) contextValidateLastUpdatedTime(ctx context.Context, formats strfmt.Registry) error {

	if err := m.LastUpdatedTime.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("lastUpdatedTime")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("lastUpdatedTime")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1ComputeMetrics) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ComputeMetrics) UnmarshalBinary(b []byte) error {
	var res V1ComputeMetrics
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
