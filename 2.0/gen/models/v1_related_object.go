// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1RelatedObject The object for which the event is related
//
// swagger:model v1RelatedObject
type V1RelatedObject struct {

	// kind
	// Enum: [spectrocluster machine cloudconfig clusterprofile pack]
	Kind string `json:"kind,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// uid
	UID string `json:"uid,omitempty"`
}

// Validate validates this v1 related object
func (m *V1RelatedObject) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateKind(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var v1RelatedObjectTypeKindPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["spectrocluster","machine","cloudconfig","clusterprofile","pack"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1RelatedObjectTypeKindPropEnum = append(v1RelatedObjectTypeKindPropEnum, v)
	}
}

const (

	// V1RelatedObjectKindSpectrocluster captures enum value "spectrocluster"
	V1RelatedObjectKindSpectrocluster string = "spectrocluster"

	// V1RelatedObjectKindMachine captures enum value "machine"
	V1RelatedObjectKindMachine string = "machine"

	// V1RelatedObjectKindCloudconfig captures enum value "cloudconfig"
	V1RelatedObjectKindCloudconfig string = "cloudconfig"

	// V1RelatedObjectKindClusterprofile captures enum value "clusterprofile"
	V1RelatedObjectKindClusterprofile string = "clusterprofile"

	// V1RelatedObjectKindPack captures enum value "pack"
	V1RelatedObjectKindPack string = "pack"
)

// prop value enum
func (m *V1RelatedObject) validateKindEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, v1RelatedObjectTypeKindPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *V1RelatedObject) validateKind(formats strfmt.Registry) error {

	if swag.IsZero(m.Kind) { // not required
		return nil
	}

	// value enum
	if err := m.validateKindEnum("kind", "body", m.Kind); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1RelatedObject) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1RelatedObject) UnmarshalBinary(b []byte) error {
	var res V1RelatedObject
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}