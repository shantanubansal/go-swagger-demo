// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1CloudConfigMeta v1 cloud config meta
//
// swagger:model v1CloudConfigMeta
type V1CloudConfigMeta struct {

	// cloud type
	CloudType V1CloudType `json:"cloudType,omitempty"`

	// Machine pool meta information
	MachinePools []*V1MachinePoolMeta `json:"machinePools"`
}

// Validate validates this v1 cloud config meta
func (m *V1CloudConfigMeta) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCloudType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMachinePools(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1CloudConfigMeta) validateCloudType(formats strfmt.Registry) error {

	if swag.IsZero(m.CloudType) { // not required
		return nil
	}

	if err := m.CloudType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("cloudType")
		}
		return err
	}

	return nil
}

func (m *V1CloudConfigMeta) validateMachinePools(formats strfmt.Registry) error {

	if swag.IsZero(m.MachinePools) { // not required
		return nil
	}

	for i := 0; i < len(m.MachinePools); i++ {
		if swag.IsZero(m.MachinePools[i]) { // not required
			continue
		}

		if m.MachinePools[i] != nil {
			if err := m.MachinePools[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("machinePools" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1CloudConfigMeta) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1CloudConfigMeta) UnmarshalBinary(b []byte) error {
	var res V1CloudConfigMeta
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
