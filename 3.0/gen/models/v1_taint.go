// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1Taint Taint
//
// swagger:model v1Taint
type V1Taint struct {

	// effect
	// Enum: [NoSchedule PreferNoSchedule NoExecute]
	Effect string `json:"effect,omitempty"`

	// The taint key to be applied to a node
	Key string `json:"key,omitempty"`

	// time added
	// Format: date-time
	TimeAdded V1Time `json:"timeAdded,omitempty"`

	// The taint value corresponding to the taint key.
	Value string `json:"value,omitempty"`
}

// Validate validates this v1 taint
func (m *V1Taint) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEffect(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimeAdded(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var v1TaintTypeEffectPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["NoSchedule","PreferNoSchedule","NoExecute"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1TaintTypeEffectPropEnum = append(v1TaintTypeEffectPropEnum, v)
	}
}

const (

	// V1TaintEffectNoSchedule captures enum value "NoSchedule"
	V1TaintEffectNoSchedule string = "NoSchedule"

	// V1TaintEffectPreferNoSchedule captures enum value "PreferNoSchedule"
	V1TaintEffectPreferNoSchedule string = "PreferNoSchedule"

	// V1TaintEffectNoExecute captures enum value "NoExecute"
	V1TaintEffectNoExecute string = "NoExecute"
)

// prop value enum
func (m *V1Taint) validateEffectEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, v1TaintTypeEffectPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *V1Taint) validateEffect(formats strfmt.Registry) error {
	if swag.IsZero(m.Effect) { // not required
		return nil
	}

	// value enum
	if err := m.validateEffectEnum("effect", "body", m.Effect); err != nil {
		return err
	}

	return nil
}

func (m *V1Taint) validateTimeAdded(formats strfmt.Registry) error {
	if swag.IsZero(m.TimeAdded) { // not required
		return nil
	}

	if err := m.TimeAdded.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("timeAdded")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("timeAdded")
		}
		return err
	}

	return nil
}

// ContextValidate validate this v1 taint based on the context it is used
func (m *V1Taint) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTimeAdded(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1Taint) contextValidateTimeAdded(ctx context.Context, formats strfmt.Registry) error {

	if err := m.TimeAdded.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("timeAdded")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("timeAdded")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1Taint) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1Taint) UnmarshalBinary(b []byte) error {
	var res V1Taint
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
