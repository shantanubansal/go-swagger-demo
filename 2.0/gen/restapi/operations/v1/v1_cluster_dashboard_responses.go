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

// V1ClusterDashboardOKCode is the HTTP code returned for type V1ClusterDashboardOK
const V1ClusterDashboardOKCode int = 200

/*V1ClusterDashboardOK An spectro cluster summary

swagger:response v1ClusterDashboardOK
*/
type V1ClusterDashboardOK struct {

	/*
	  In: Body
	*/
	Payload *models.V1ClusterSummary `json:"body,omitempty"`
}

// NewV1ClusterDashboardOK creates V1ClusterDashboardOK with default headers values
func NewV1ClusterDashboardOK() *V1ClusterDashboardOK {

	return &V1ClusterDashboardOK{}
}

// WithPayload adds the payload to the v1 cluster dashboard o k response
func (o *V1ClusterDashboardOK) WithPayload(payload *models.V1ClusterSummary) *V1ClusterDashboardOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v1 cluster dashboard o k response
func (o *V1ClusterDashboardOK) SetPayload(payload *models.V1ClusterSummary) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V1ClusterDashboardOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
