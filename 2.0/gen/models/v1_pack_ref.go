// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1PackRef PackRef server/name:tag to point to a pack PackRef is used when construct a ClusterProfile PackSpec is used for UI to render the parameters form ClusterProfile will not know inner details of a pack ClusterProfile only contain pack name:tag, and the param values user entered for it
//
// swagger:model v1PackRef
type V1PackRef struct {

	// Annotations is used to allow packref to add more arbitrary information one example is to add git reference for values.yaml
	Annotations map[string]string `json:"annotations,omitempty"`

	// digest is used to specify the version should be installed by palette when pack upgrade available, change this digest to trigger upgrade
	Digest string `json:"digest,omitempty"`

	// in valid reason
	InValidReason string `json:"inValidReason,omitempty"`

	// pack is invalid when the associated tag is deleted from the registry
	IsInvalid bool `json:"isInvalid,omitempty"`

	// layer
	// Required: true
	// Enum: [kernel os k8s cni csi addon]
	Layer *string `json:"layer"`

	// manifests
	Manifests []*V1ObjectReference `json:"manifests"`

	// pack name
	// Required: true
	Name *string `json:"name"`

	// PackUID is Hubble packUID, not palette Pack.UID It is used by Hubble only.
	PackUID string `json:"packUid,omitempty"`

	// params passed as env variables to be consumed at installation time
	Params map[string]string `json:"params,omitempty"`

	// presets
	Presets []*V1PackPreset `json:"presets"`

	// pack registry uid
	RegistryUID string `json:"registryUid,omitempty"`

	// schema
	Schema []*V1PackSchema `json:"schema"`

	// pack registry server or helm repo
	Server string `json:"server,omitempty"`

	// pack tag
	Tag string `json:"tag,omitempty"`

	// type of the pack
	// Enum: [spectro helm manifest]
	Type string `json:"type,omitempty"`

	// values represents the values.yaml used as input parameters either Params OR Values should be used, not both If both applied at the same time, will only use Values
	Values string `json:"values,omitempty"`

	// pack version
	Version string `json:"version,omitempty"`
}

// Validate validates this v1 pack ref
func (m *V1PackRef) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLayer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateManifests(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePresets(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSchema(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var v1PackRefTypeLayerPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["kernel","os","k8s","cni","csi","addon"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1PackRefTypeLayerPropEnum = append(v1PackRefTypeLayerPropEnum, v)
	}
}

const (

	// V1PackRefLayerKernel captures enum value "kernel"
	V1PackRefLayerKernel string = "kernel"

	// V1PackRefLayerOs captures enum value "os"
	V1PackRefLayerOs string = "os"

	// V1PackRefLayerK8s captures enum value "k8s"
	V1PackRefLayerK8s string = "k8s"

	// V1PackRefLayerCni captures enum value "cni"
	V1PackRefLayerCni string = "cni"

	// V1PackRefLayerCsi captures enum value "csi"
	V1PackRefLayerCsi string = "csi"

	// V1PackRefLayerAddon captures enum value "addon"
	V1PackRefLayerAddon string = "addon"
)

// prop value enum
func (m *V1PackRef) validateLayerEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, v1PackRefTypeLayerPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *V1PackRef) validateLayer(formats strfmt.Registry) error {

	if err := validate.Required("layer", "body", m.Layer); err != nil {
		return err
	}

	// value enum
	if err := m.validateLayerEnum("layer", "body", *m.Layer); err != nil {
		return err
	}

	return nil
}

func (m *V1PackRef) validateManifests(formats strfmt.Registry) error {

	if swag.IsZero(m.Manifests) { // not required
		return nil
	}

	for i := 0; i < len(m.Manifests); i++ {
		if swag.IsZero(m.Manifests[i]) { // not required
			continue
		}

		if m.Manifests[i] != nil {
			if err := m.Manifests[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("manifests" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1PackRef) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *V1PackRef) validatePresets(formats strfmt.Registry) error {

	if swag.IsZero(m.Presets) { // not required
		return nil
	}

	for i := 0; i < len(m.Presets); i++ {
		if swag.IsZero(m.Presets[i]) { // not required
			continue
		}

		if m.Presets[i] != nil {
			if err := m.Presets[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("presets" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1PackRef) validateSchema(formats strfmt.Registry) error {

	if swag.IsZero(m.Schema) { // not required
		return nil
	}

	for i := 0; i < len(m.Schema); i++ {
		if swag.IsZero(m.Schema[i]) { // not required
			continue
		}

		if m.Schema[i] != nil {
			if err := m.Schema[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("schema" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var v1PackRefTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["spectro","helm","manifest"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1PackRefTypeTypePropEnum = append(v1PackRefTypeTypePropEnum, v)
	}
}

const (

	// V1PackRefTypeSpectro captures enum value "spectro"
	V1PackRefTypeSpectro string = "spectro"

	// V1PackRefTypeHelm captures enum value "helm"
	V1PackRefTypeHelm string = "helm"

	// V1PackRefTypeManifest captures enum value "manifest"
	V1PackRefTypeManifest string = "manifest"
)

// prop value enum
func (m *V1PackRef) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, v1PackRefTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *V1PackRef) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1PackRef) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1PackRef) UnmarshalBinary(b []byte) error {
	var res V1PackRef
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
