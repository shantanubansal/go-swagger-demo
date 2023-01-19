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

// NewV1ClusterDashboardParams creates a new V1ClusterDashboardParams object
//
// There are no default values defined in the spec.
func NewV1ClusterDashboardParams() V1ClusterDashboardParams {

	return V1ClusterDashboardParams{}
}

// V1ClusterDashboardParams contains all the bound params for the v1 cluster dashboard operation
// typically these are obtained from a http.Request
//
// swagger:parameters v1ClusterDashboard
type V1ClusterDashboardParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewV1ClusterDashboardParams() beforehand.
func (o *V1ClusterDashboardParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
