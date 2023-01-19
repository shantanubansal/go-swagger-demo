// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/shantanubansal/go-swagger-demo/3.0/gen/restapi/operations"
	"github.com/shantanubansal/go-swagger-demo/3.0/gen/restapi/operations/v1"
)

//go:generate swagger generate server --target ../../gen --name SwagAPI --spec ../../../3.0/spec/swag_api.yaml --principal interface{}

func configureFlags(api *operations.SwagAPIAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.SwagAPIAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	if api.V1V1SystemHealthPingHandler == nil {
		api.V1V1SystemHealthPingHandler = v1.V1SystemHealthPingHandlerFunc(func(params v1.V1SystemHealthPingParams) middleware.Responder {
			return middleware.NotImplemented("operation v1.V1SystemHealthPing has not yet been implemented")
		})
	}
	if api.V1V1AuthenticateHandler == nil {
		api.V1V1AuthenticateHandler = v1.V1AuthenticateHandlerFunc(func(params v1.V1AuthenticateParams) middleware.Responder {
			return middleware.NotImplemented("operation v1.V1Authenticate has not yet been implemented")
		})
	}
	if api.V1V1CloudConfigAwsHandler == nil {
		api.V1V1CloudConfigAwsHandler = v1.V1CloudConfigAwsHandlerFunc(func(params v1.V1CloudConfigAwsParams) middleware.Responder {
			return middleware.NotImplemented("operation v1.V1CloudConfigAws has not yet been implemented")
		})
	}
	if api.V1V1ClusterDashboardHandler == nil {
		api.V1V1ClusterDashboardHandler = v1.V1ClusterDashboardHandlerFunc(func(params v1.V1ClusterDashboardParams) middleware.Responder {
			return middleware.NotImplemented("operation v1.V1ClusterDashboard has not yet been implemented")
		})
	}
	if api.V1V1ClusterEventsHandler == nil {
		api.V1V1ClusterEventsHandler = v1.V1ClusterEventsHandlerFunc(func(params v1.V1ClusterEventsParams) middleware.Responder {
			return middleware.NotImplemented("operation v1.V1ClusterEvents has not yet been implemented")
		})
	}
	if api.V1V1ClusterOverviewHandler == nil {
		api.V1V1ClusterOverviewHandler = v1.V1ClusterOverviewHandlerFunc(func(params v1.V1ClusterOverviewParams) middleware.Responder {
			return middleware.NotImplemented("operation v1.V1ClusterOverview has not yet been implemented")
		})
	}
	if api.V1V1UserInfoHandler == nil {
		api.V1V1UserInfoHandler = v1.V1UserInfoHandlerFunc(func(params v1.V1UserInfoParams) middleware.Responder {
			return middleware.NotImplemented("operation v1.V1UserInfo has not yet been implemented")
		})
	}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
