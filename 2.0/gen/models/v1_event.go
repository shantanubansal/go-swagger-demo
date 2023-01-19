// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1Event Event information organized by tags
//
// swagger:model v1Event
type V1Event struct {

	// involved object
	InvolvedObject *V1ObjectReference `json:"involvedObject,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// metadata
	Metadata *V1ObjectMeta `json:"metadata,omitempty"`

	// reason
	Reason string `json:"reason,omitempty"`

	// related object
	RelatedObject *V1RelatedObject `json:"relatedObject,omitempty"`

	// severity
	Severity string `json:"severity,omitempty"`

	// source
	Source *V1EventSource `json:"source,omitempty"`
}

// Validate validates this v1 event
func (m *V1Event) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInvolvedObject(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetadata(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRelatedObject(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSource(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1Event) validateInvolvedObject(formats strfmt.Registry) error {

	if swag.IsZero(m.InvolvedObject) { // not required
		return nil
	}

	if m.InvolvedObject != nil {
		if err := m.InvolvedObject.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("involvedObject")
			}
			return err
		}
	}

	return nil
}

func (m *V1Event) validateMetadata(formats strfmt.Registry) error {

	if swag.IsZero(m.Metadata) { // not required
		return nil
	}

	if m.Metadata != nil {
		if err := m.Metadata.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metadata")
			}
			return err
		}
	}

	return nil
}

func (m *V1Event) validateRelatedObject(formats strfmt.Registry) error {

	if swag.IsZero(m.RelatedObject) { // not required
		return nil
	}

	if m.RelatedObject != nil {
		if err := m.RelatedObject.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("relatedObject")
			}
			return err
		}
	}

	return nil
}

func (m *V1Event) validateSource(formats strfmt.Registry) error {

	if swag.IsZero(m.Source) { // not required
		return nil
	}

	if m.Source != nil {
		if err := m.Source.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("source")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1Event) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1Event) UnmarshalBinary(b []byte) error {
	var res V1Event
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}