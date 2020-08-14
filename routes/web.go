package routes

import (
    "github.com/gorilla/mux"
    "log"
    "fmt"
    "net/http"
    "github.com/LayssonENS/go-desafio/Controllers" 
)

func ConfigurarRotas()  {

    router := mux.NewRouter()
    router.HandleFunc("/cnpj/{cnpj}", Controllers.BuscarCnpj).Methods("GET")
    router.HandleFunc("/posts", Controllers.BuscarFrete).Methods("POST")

    fmt.Println("Servdor executando")
    log.Fatal(http.ListenAndServe(":3000", router))

}



