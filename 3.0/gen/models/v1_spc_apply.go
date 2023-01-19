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

// V1SpcApply v1 spc apply
//
// swagger:model v1SpcApply
type V1SpcApply struct {

	// action type
	// Enum: [DownloadAndInstall DownloadAndInstallLater]
	ActionType string `json:"actionType,omitempty"`

	// If it is true then Agent can apply the changes to the palette
	CanBeApplied bool `json:"canBeApplied"`

	// last modified time
	// Format: date-time
	LastModifiedTime V1Time `json:"lastModifiedTime,omitempty"`

	// patch applied time
	// Format: date-time
	PatchAppliedTime V1Time `json:"patchAppliedTime,omitempty"`
}

// Validate validates this v1 spc apply
func (m *V1SpcApply) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActionType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastModifiedTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePatchAppliedTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var v1SpcApplyTypeActionTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["DownloadAndInstall","DownloadAndInstallLater"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1SpcApplyTypeActionTypePropEnum = append(v1SpcApplyTypeActionTypePropEnum, v)
	}
}

const (

	// V1SpcApplyActionTypeDownloadAndInstall captures enum value "DownloadAndInstall"
	V1SpcApplyActionTypeDownloadAndInstall string = "DownloadAndInstall"

	// V1SpcApplyActionTypeDownloadAndInstallLater captures enum value "DownloadAndInstallLater"
	V1SpcApplyActionTypeDownloadAndInstallLater string = "DownloadAndInstallLater"
)

// prop value enum
func (m *V1SpcApply) validateActionTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, v1SpcApplyTypeActionTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *V1SpcApply) validateActionType(formats strfmt.Registry) error {
	if swag.IsZero(m.ActionType) { // not required
		return nil
	}

	// value enum
	if err := m.validateActionTypeEnum("actionType", "body", m.ActionType); err != nil {
		return err
	}

	return nil
}

func (m *V1SpcApply) validateLastModifiedTime(formats strfmt.Registry) error {
	if swag.IsZero(m.LastModifiedTime) { // not required
		return nil
	}

	if err := m.LastModifiedTime.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("lastModifiedTime")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("lastModifiedTime")
		}
		return err
	}

	return nil
}

func (m *V1SpcApply) validatePatchAppliedTime(formats strfmt.Registry) error {
	if swag.IsZero(m.PatchAppliedTime) { // not required
		return nil
	}

	if err := m.PatchAppliedTime.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("patchAppliedTime")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("patchAppliedTime")
		}
		return err
	}

	return nil
}

// ContextValidate validate this v1 spc apply based on the context it is used
func (m *V1SpcApply) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLastModifiedTime(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePatchAppliedTime(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1SpcApply) contextValidateLastModifiedTime(ctx context.Context, formats strfmt.Registry) error {

	if err := m.LastModifiedTime.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("lastModifiedTime")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("lastModifiedTime")
		}
		return err
	}

	return nil
}

func (m *V1SpcApply) contextValidatePatchAppliedTime(ctx context.Context, formats strfmt.Registry) error {

	if err := m.PatchAppliedTime.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("patchAppliedTime")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("patchAppliedTime")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1SpcApply) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1SpcApply) UnmarshalBinary(b []byte) error {
	var res V1SpcApply
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
