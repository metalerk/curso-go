package main

import (

	"net/http"
	"io"

)

func main(){
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request){
	io.WriteString(w, "<h1>Hola Mundo !!</h1>")
}
