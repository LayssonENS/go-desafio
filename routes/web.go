package routes

import (
	"fmt"
	"log"
	"net/http"

	"github.com/LayssonENS/go-desafio/controllers"
	"github.com/gorilla/mux"
)

/*ConfigurarRotas endPoint*/
func ConfigurarRotas() {
	router := mux.NewRouter()
	router.HandleFunc("/cnpj/{cnpj}", controllers.BuscarCnpj).Methods("GET")
	router.HandleFunc("/quote", controllers.BuscarFrete).Methods("POST")

	fmt.Println("Servdor executando")
	log.Fatal(http.ListenAndServe(":3000", router))

}
