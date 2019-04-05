// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"log"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/mbobakov/showcase/restapi/operations"
	"github.com/mbobakov/showcase/restapi/operations/metrics"
	"github.com/mbobakov/showcase/service/metric"
	"github.com/mbobakov/showcase/storage/clickhouse"
)

//go:generate swagger generate server --target ../../showcase --name AClickhouseSwaggerApplication --spec ../swagger.yml

var chConfig = struct {
	DSN   string `long:"--ch-dsn" default:"tcp://127.0.0.1:9000" description:"Clickhouse connection string"`
	Table string `long:"--ch-table" default:"ts" description:"Clickhouse table for the data"`
}{}

func configureFlags(api *operations.AClickhouseSwaggerApplicationAPI) {
	api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{
		swag.CommandLineOptionsGroup{
			LongDescription:  "Clickhouse options",
			ShortDescription: "Clickhouse options",
			Options:          &chConfig,
		},
	}
}

func configureAPI(api *operations.AClickhouseSwaggerApplicationAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()
	api.JSONProducer = runtime.JSONProducer()

	db, err := clickhouse.Init(chConfig.DSN, chConfig.Table)
	if err != nil {
		log.Fatal(err)
	}
	srvc := metric.New(db)

	api.MetricsFindMetricsHandler = metrics.FindMetricsHandlerFunc(srvc.FindMetrics)
	api.MetricsPostDatapointHandler = metrics.PostDatapointHandlerFunc(srvc.PostDatapoint)

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
