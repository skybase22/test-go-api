package router

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	//"strconv"
)

type Server struct {
	router http.Handler
	port   int
}

// Router router
type Router struct {
	router *mux.Router
}

type addressBook struct {
	Firstname string `json:"first_name"`
	Lastname  string `json:"last_name"`
	Code      int    `json:"code,omitempty"`
	Phone     string `json:"phone,omitempty"`
}

func New() {
	router := mux.NewRouter()

	api := router.PathPrefix("/api").Subrouter().StrictSlash(true)
	version := api.PathPrefix("/{version}").Subrouter().StrictSlash(true)

	guest := version.PathPrefix("/guest").Subrouter().StrictSlash(true)

	guest.HandleFunc("/home", getData()).Methods("GET")
	guest.HandleFunc("/name", sendData()).Methods("POST")

	if err := http.ListenAndServe(":8000", router); err != nil {
		logrus.Panic("http.ListenAndServe: ", err)
	}

	//return &Server{router, port}

}

func getData() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		addNew := &addressBook{
			Firstname: "Nop",
			Lastname:  "San",
		}
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Cache-Control", "max-age=0")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(addNew)
	}
}

func sendData() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		addBook := &addressBook{}
		err := json.NewDecoder(r.Body).Decode(addBook)
		if err != nil {
			return
		}
		fmt.Println(addBook)
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Cache-Control", "max-age=0")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(addBook)
	}
}

// func (s *Server) Start() {
// 	portString := fmt.Sprintf(":%d", s.port)
// 	logrus.Infof("Listening on port %s", portString)
// 	fmt.Println(s.port, s.router)
// 	if err := http.ListenAndServe(portString, s.router); err != nil {
// 		logrus.Panic("http.ListenAndServe: ", err)
// 	}
// }
