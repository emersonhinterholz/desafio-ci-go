package main

import (
	"fmt"
	"log"
	"net/http"

	"../utils"
)

func main() {
	http.HandleFunc("/", helloServer)
	log.Fatal(http.ListenAndServe(":80", nil))
}

func helloServer(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("Code Rocks! 2 + 2 = %d", utils.Sum(2, 2))))
}
