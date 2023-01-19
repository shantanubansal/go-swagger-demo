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

// V1ClusterProfileTemplate ClusterProfileTemplate contains details of a clusterprofile definition
//
// swagger:model v1ClusterProfileTemplate
type V1ClusterProfileTemplate struct {

	// cloud type
	CloudType V1CloudType `json:"cloudType,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// PackServerRefs is only used on Hubble side it is reference to pack registry servers which PackRef belongs to in hubble, pack server is a top level object, so use a reference to point to it packs within a clusterprofile can come from different pack servers, so this is an array
	PackServerRefs []*V1ObjectReference `json:"packServerRefs"`

	// This secret is used only on Palette side use case is similar  to k8s image pull secret this single secret internally should contains all the pack servers in PackServerRefs if empty, means no credential is needed to access the pack server For spectro saas, Ally will set this field before pass to palette
	PackServerSecret string `json:"packServerSecret,omitempty"`

	// Packs definitions here are final definitions. If ClonedFrom and ParamsOverwrite is present, then Packs are the final merge result of ClonedFrom and ParamsOverwrite So orchestration engine will just take the Packs and do the work, no need to worry about parameters merge
	Packs []*V1PackRef `json:"packs"`

	// version start from 1.0.0, matching the index of ClusterProfileSpec.Versions[] will be used by clusterSpec to identify which version is applied to the cluster
	ProfileVersion string `json:"profileVersion,omitempty"`

	// type
	Type string `json:"type,omitempty"`

	// uid
	UID string `json:"uid,omitempty"`

	// Deprecated. Use profileVersion
	Version int32 `json:"version,omitempty"`
}

// Validate validates this v1 cluster profile template
func (m *V1ClusterProfileTemplate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCloudType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePackServerRefs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePacks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ClusterProfileTemplate) validateCloudType(formats strfmt.Registry) error {

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

func (m *V1ClusterProfileTemplate) validatePackServerRefs(formats strfmt.Registry) error {

	if swag.IsZero(m.PackServerRefs) { // not required
		return nil
	}

	for i := 0; i < len(m.PackServerRefs); i++ {
		if swag.IsZero(m.PackServerRefs[i]) { // not required
			continue
		}

		if m.PackServerRefs[i] != nil {
			if err := m.PackServerRefs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("packServerRefs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1ClusterProfileTemplate) validatePacks(formats strfmt.Registry) error {

	if swag.IsZero(m.Packs) { // not required
		return nil
	}

	for i := 0; i < len(m.Packs); i++ {
		if swag.IsZero(m.Packs[i]) { // not required
			continue
		}

		if m.Packs[i] != nil {
			if err := m.Packs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("packs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1ClusterProfileTemplate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ClusterProfileTemplate) UnmarshalBinary(b []byte) error {
	var res V1ClusterProfileTemplate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
