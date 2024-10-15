package main

import(
	"encoding/json"
	"log"
	"net/http"
	"goserver/model"
)

func handler (w http.ResponseWriter,  r *http.Request){
	response :=model.Response{Message:"hello depuis le serveur Go"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main(){
	http.HandleFunc("/api/greet", handler)
    log.Println("Go server running on http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}