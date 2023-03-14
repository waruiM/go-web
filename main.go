package main

import (
	"fmt"
	"log"
	"net/http"

)


func st1func(w http.ResponseWriter,r *http.Request ){
    //where the url is wrong
	if r.URL.Path != "/st1"{
		http.Error(w, "404 page not there",http.StatusNotFound)
		return
	}

	fmt.Fprintf(w,"page st1")
}
func st2func(w http.ResponseWriter,r *http.Request){
	if r.URL.Path != "/st2"{
		http.Error(w,"401 chickens missing",http.StatusNotFound)
		return
	}

	fmt.Fprintf(w,"page st2")
}




func main(){
	FileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", FileServer)
	http.HandleFunc("/st1",st1func)
	http.HandleFunc("/st2",st2func)

	fmt.Printf("Starting server at port 8080")

	if err :=http.ListenAndServe(":8080",nil);err!=nil{
		log.Fatal(err)
	}

}