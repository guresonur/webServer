package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	http.HandleFunc("/", handler)
	if len(os.Args) > 1 {
	log.Fatal(http.ListenAndServe("localhost:"+os.Args[1], nil))
}	else {
		log.Fatal(http.ListenAndServe("localhost:8090", nil))
	}

}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "running at port:%q", os.Args[1])
}
