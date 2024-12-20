// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"database/sql"
	"log"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"backend-swagger/handlers"
	"backend-swagger/postgres"
	"backend-swagger/restapi/operations"
	"backend-swagger/utils"
)

func postgres_connection() *sql.DB {
	db, err := postgres.ConnectDB()
	if err != nil {
		log.Fatal("Error al conectar con la base de datos: ", err)
	}
	return db
}

//go:generate swagger generate server --target ..\..\backend-swagger --name UserManagementAPI --spec ..\swagger.yaml --principal interface{}

// In Windows inside this folder backend-swagger use:
// swagger generate server --target . --name UserManagementAPI --spec "swagger.yaml" --principal "interface{}"
func configureFlags(api *operations.UserManagementAPIAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.UserManagementAPIAPI) http.Handler {
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

	api.BearerAuthAuth = utils.BearerTokenAuth

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	newdb := postgres_connection()

	api.DeleteUsersIDHandler = operations.DeleteUsersIDHandlerFunc(func(params operations.DeleteUsersIDParams, principal interface{}) middleware.Responder {
		handler := handlers.NewUserHandler(newdb)
		return handler.DeleteUser(params)
	})

	api.GetUsersHandler = operations.GetUsersHandlerFunc(func(params operations.GetUsersParams, principal interface{}) middleware.Responder {
		handler := handlers.NewUserHandler(newdb)
		return handler.GetUsers(params)
	})

	api.PostUsersHandler = operations.PostUsersHandlerFunc(func(params operations.PostUsersParams, principal interface{}) middleware.Responder {
		handler := handlers.NewUserHandler(newdb)
		return handler.AddUser(params)
	})

	api.PutUsersIDHandler = operations.PutUsersIDHandlerFunc(func(params operations.PutUsersIDParams, principal interface{}) middleware.Responder {
		handler := handlers.NewUserHandler(newdb)
		return handler.UpdateUser(params)
	})

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
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
