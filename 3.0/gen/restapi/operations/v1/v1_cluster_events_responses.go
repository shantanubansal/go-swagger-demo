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

// V1ClusterEventsOKCode is the HTTP code returned for type V1ClusterEventsOK
const V1ClusterEventsOKCode int = 200

/*
V1ClusterEventsOK An array of component event items

swagger:response v1ClusterEventsOK
*/
type V1ClusterEventsOK struct {

	/*
	  In: Body
	*/
	Payload *models.V1ClusterEvents `json:"body,omitempty"`
}

// NewV1ClusterEventsOK creates V1ClusterEventsOK with default headers values
func NewV1ClusterEventsOK() *V1ClusterEventsOK {

	return &V1ClusterEventsOK{}
}

// WithPayload adds the payload to the v1 cluster events o k response
func (o *V1ClusterEventsOK) WithPayload(payload *models.V1ClusterEvents) *V1ClusterEventsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v1 cluster events o k response
func (o *V1ClusterEventsOK) SetPayload(payload *models.V1ClusterEvents) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V1ClusterEventsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
