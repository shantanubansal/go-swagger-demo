// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1EksClusterConfig EksClusterConfig defines EKS specific config
//
// swagger:model v1EksClusterConfig
type V1EksClusterConfig struct {

	// Addons defines the EKS addons to enable with the EKS cluster. This may be required for brownfield clusters
	Addons []*V1EksAddon `json:"addons"`

	// BastionDisabled is the option to disable bastion node
	BastionDisabled bool `json:"bastionDisabled,omitempty"`

	// ControlPlaneLoadBalancer specifies how API server elb will be configured, this field is optional, not provided, "", default => "Internet-facing" "Internet-facing" => "Internet-facing" "internal" => "internal" For spectro saas setup we require to talk to the apiserver from our cluster so ControlPlaneLoadBalancer should be "", not provided or "Internet-facing"
	ControlPlaneLoadBalancer string `json:"controlPlaneLoadBalancer,omitempty"`

	// encryption config
	EncryptionConfig *V1EncryptionConfig `json:"encryptionConfig,omitempty"`

	// endpoint access
	EndpointAccess *V1EksClusterConfigEndpointAccess `json:"endpointAccess,omitempty"`

	// The AWS Region the cluster lives in.
	// Required: true
	Region *string `json:"region"`

	// SSHKeyName specifies which EC2 SSH key can be used to access machines.
	SSHKeyName string `json:"sshKeyName,omitempty"`

	// VPC Id to deploy cluster into should have one public and one private subnet for the the cluster creation, this field is optional, If VPC Id is not provided a fully managed VPC will be created
	VpcID string `json:"vpcId,omitempty"`
}

// Validate validates this v1 eks cluster config
func (m *V1EksClusterConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddons(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEncryptionConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndpointAccess(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1EksClusterConfig) validateAddons(formats strfmt.Registry) error {

	if swag.IsZero(m.Addons) { // not required
		return nil
	}

	for i := 0; i < len(m.Addons); i++ {
		if swag.IsZero(m.Addons[i]) { // not required
			continue
		}

		if m.Addons[i] != nil {
			if err := m.Addons[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("addons" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1EksClusterConfig) validateEncryptionConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.EncryptionConfig) { // not required
		return nil
	}

	if m.EncryptionConfig != nil {
		if err := m.EncryptionConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("encryptionConfig")
			}
			return err
		}
	}

	return nil
}

func (m *V1EksClusterConfig) validateEndpointAccess(formats strfmt.Registry) error {

	if swag.IsZero(m.EndpointAccess) { // not required
		return nil
	}

	if m.EndpointAccess != nil {
		if err := m.EndpointAccess.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("endpointAccess")
			}
			return err
		}
	}

	return nil
}

func (m *V1EksClusterConfig) validateRegion(formats strfmt.Registry) error {

	if err := validate.Required("region", "body", m.Region); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1EksClusterConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1EksClusterConfig) UnmarshalBinary(b []byte) error {
	var res V1EksClusterConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
