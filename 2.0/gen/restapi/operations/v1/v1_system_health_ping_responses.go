// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/shantanubansal/go-swagger-demo/3.0/gen/models"
	"github.com/shantanubansal/go-swagger-demo/3.0/gen/restapi/operations/v1"
)

// V1SystemHealthPingOKCode is the HTTP code returned for type V1SystemHealthPingOK
const V1SystemHealthPingOKCode int = 200

/*V1SystemHealthPingOK OK

swagger:response v1SystemHealthPingOK
*/
type V1SystemHealthPingOK struct {

	/*
	  In: Body
	*/
	Payload *models.V1SystemHealthPing `json:"body,omitempty"`
}

// NewV1SystemHealthPingOK creates V1SystemHealthPingOK with default headers values
func NewV1SystemHealthPingOK() *V1SystemHealthPingOK {

	return &V1SystemHealthPingOK{}
}

// WithPayload adds the payload to the v1 system health ping o k response
func (o *V1SystemHealthPingOK) WithPayload(payload *models.V1SystemHealthPing) *V1SystemHealthPingOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v1 system health ping o k response
func (o *V1SystemHealthPingOK) SetPayload(payload *models.V1SystemHealthPing) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V1SystemHealthPingOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
