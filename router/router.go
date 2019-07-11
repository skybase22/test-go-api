package router

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Server struct {
	router http.Handler
	port   int
}

// Router router
type Router struct {
	router *mux.Router
}

func New(port int) *Server {
	router := mux.NewRouter()

	api := router.PathPrefix("/api").Subrouter().StrictSlash(true)
	version := api.PathPrefix("/{version}").Subrouter().StrictSlash(true)

	guest := version.PathPrefix("/guest").Subrouter().StrictSlash(true)

	guest.HandleFunc("/register", getData()).Methods("GET")

	return &Server{router, port}

}

func getData() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Welcome to the HomePage!")
	}
}

func (s *Server) Start() {
	portString := fmt.Sprintf(":%d", s.port)
	logrus.Infof("Listening on port %s", portString)
	if err := http.ListenAndServe(portString, s.router); err != nil {
		logrus.Panic("http.ListenAndServe: ", err)
	}
}
