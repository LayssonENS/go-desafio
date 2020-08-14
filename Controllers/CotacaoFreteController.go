package Controllers

import (
   // "github.com/LayssonENS/go-desafio/Models"
    "encoding/json"
    "net/http"
   // "io/ioutil"
    "github.com/gorilla/mux"
) 

func BuscarFrete(w http.ResponseWriter, r *http.Request) {
    
    /*Pegando parâmetro passado via GET*/
	dadosRequest := mux.Vars(r)

    /*Retorna dados personalizados*/
    json.NewEncoder(w).Encode(dadosRequest)
    
}
