// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"

	"github.com/shantanubansal/go-swagger-demo/3.0/gen/restapi/operations/v1"
)

// NewV1UserInfoParams creates a new V1UserInfoParams object
// no default values defined in spec.
func NewV1UserInfoParams() V1UserInfoParams {

	return V1UserInfoParams{}
}

// V1UserInfoParams contains all the bound params for the v1 user info operation
// typically these are obtained from a http.Request
//
// swagger:parameters v1UserInfo
type V1UserInfoParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewV1UserInfoParams() beforehand.
func (o *V1UserInfoParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
