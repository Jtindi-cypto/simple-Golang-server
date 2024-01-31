package main

import (
"fmt"
"log"
"net/http"
)

func handler(w http.ResponseWriter, r *http.Request){
     fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
}

func main(){
http.HandleFunc("/", handler)

fmt.Println("The server is running on http://localhost:8080")
http.ListenAndServe(":8080",nil)

}
