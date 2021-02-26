package app

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
	"github.com/vipul-08/cassandra-api/domain"
	"github.com/vipul-08/cassandra-api/service"
	"log"
	"net/http"
	"os"
)

func StartRoutes() {
	router := mux.NewRouter()

	handler := StudentHandlers{service.NewStudentService(domain.NewStudentRepository())}

	router.HandleFunc("/students", handler.getAllStudents).Methods(http.MethodGet)
	router.HandleFunc("/students/{id}", handler.getStudentById).Methods(http.MethodGet)
	router.HandleFunc("/students", handler.insertStudent).Methods(http.MethodPost)
	router.HandleFunc("/students/{id}", handler.deleteStudent).Methods(http.MethodDelete)
	router.HandleFunc("/students/{id}", handler.updateStudent).Methods(http.MethodPut)

	ops := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.Redoc(ops, nil)
	router.Handle("/docs", sh)
	router.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router))
}
