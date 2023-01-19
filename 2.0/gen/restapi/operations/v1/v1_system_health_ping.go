// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/shantanubansal/go-swagger-demo/3.0/gen/restapi/operations/v1"
)

// V1SystemHealthPingHandlerFunc turns a function with the right signature into a v1 system health ping handler
type V1SystemHealthPingHandlerFunc func(V1SystemHealthPingParams) middleware.Responder

// Handle executing the request and returning a response
func (fn V1SystemHealthPingHandlerFunc) Handle(params V1SystemHealthPingParams) middleware.Responder {
	return fn(params)
}

// V1SystemHealthPingHandler interface for that can handle valid v1 system health ping params
type V1SystemHealthPingHandler interface {
	Handle(V1SystemHealthPingParams) middleware.Responder
}

// NewV1SystemHealthPing creates a new http.Handler for the v1 system health ping operation
func NewV1SystemHealthPing(ctx *middleware.Context, handler V1SystemHealthPingHandler) *V1SystemHealthPing {
	return &V1SystemHealthPing{Context: ctx, Handler: handler}
}

/*V1SystemHealthPing swagger:route GET /v1/system/health/ping v1 private v1SystemHealthPing

Ping Service

Ping Service

*/
type V1SystemHealthPing struct {
	Context *middleware.Context
	Handler V1SystemHealthPingHandler
}

func (o *V1SystemHealthPing) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewV1SystemHealthPingParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
