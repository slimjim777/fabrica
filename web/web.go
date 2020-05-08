package web

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/ogra1/fabrica/config"
	"github.com/ogra1/fabrica/service"
	"net/http"
)

// Web implements the web service
type Web struct {
	Settings *config.Settings
	BuildSrv service.BuildSrv
}

// NewWebService starts a new web service
func NewWebService(settings *config.Settings, bldSrv service.BuildSrv) *Web {
	return &Web{
		Settings: settings,
		BuildSrv: bldSrv,
	}
}

// Start the web service
func (srv Web) Start() error {
	listenOn := fmt.Sprintf("%s:%s", "0.0.0.0", srv.Settings.Port)
	fmt.Printf("Starting service on port %s\n", listenOn)
	return http.ListenAndServe(listenOn, srv.Router())
}

// Router returns the application router
func (srv Web) Router() *mux.Router {
	// Start the web service router
	router := mux.NewRouter()

	router.Handle("/v1/build", Middleware(http.HandlerFunc(srv.Build))).Methods("POST")
	router.Handle("/v1/builds", Middleware(http.HandlerFunc(srv.BuildList))).Methods("GET")
	router.Handle("/v1/builds/{id}", Middleware(http.HandlerFunc(srv.BuildLog))).Methods("GET")

	// Serve the static path
	fs := http.StripPrefix("/static/", http.FileServer(http.Dir(docRoot)))
	router.PathPrefix("/static/").Handler(fs)

	// Default path is the index page
	router.Handle("/", Middleware(http.HandlerFunc(srv.Index))).Methods("GET")
	router.Handle("/builds/{id}", Middleware(http.HandlerFunc(srv.Index))).Methods("GET")

	return router
}
