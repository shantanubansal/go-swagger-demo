// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/shantanubansal/go-swagger-demo/3.0/gen/models"
	"github.com/shantanubansal/go-swagger-demo/3.0/gen/restapi/operations/v1"
)

// NewV1AuthenticateParams creates a new V1AuthenticateParams object
// with the default values initialized.
func NewV1AuthenticateParams() V1AuthenticateParams {

	var (
		// initialize parameters with default values

		setCookieDefault = bool(true)
	)

	return V1AuthenticateParams{
		SetCookie: &setCookieDefault,
	}
}

// V1AuthenticateParams contains all the bound params for the v1 authenticate operation
// typically these are obtained from a http.Request
//
// swagger:parameters v1Authenticate
type V1AuthenticateParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: body
	*/
	Body *models.V1AuthLogin
	/*
	  In: query
	  Default: true
	*/
	SetCookie *bool
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewV1AuthenticateParams() beforehand.
func (o *V1AuthenticateParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.V1AuthLogin
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("body", "body", ""))
			} else {
				res = append(res, errors.NewParseError("body", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Body = &body
			}
		}
	} else {
		res = append(res, errors.Required("body", "body", ""))
	}
	qSetCookie, qhkSetCookie, _ := qs.GetOK("setCookie")
	if err := o.bindSetCookie(qSetCookie, qhkSetCookie, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindSetCookie binds and validates parameter SetCookie from query.
func (o *V1AuthenticateParams) bindSetCookie(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewV1AuthenticateParams()
		return nil
	}

	value, err := swag.ConvertBool(raw)
	if err != nil {
		return errors.InvalidType("setCookie", "query", "bool", raw)
	}
	o.SetCookie = &value

	return nil
}
