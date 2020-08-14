package routes

import (
    "github.com/gorilla/mux"
    "Api/Controllers"
    "log"
    "fmt"
    "net/http" 
)

func ConfigurarRotas()  {

    router := mux.NewRouter()
    router.HandleFunc("/cnpj/{cnpj}", Controllers.BuscarCnpj).Methods("GET")

    fmt.Println("Servdor executando")
    log.Fatal(http.ListenAndServe(":3000", router))

}



