package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:"+os.Args[1], nil))

}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "running at port:%q", os.Args[1])
}
