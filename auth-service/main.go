package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	/*
		todo search rfid id ; if exist => jwt token to client
		todo postgres setup + rfid simulation
	*/
	http.HandleFunc("/auth", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(writer, "Auth service")
	})
	log.Fatalln(http.ListenAndServe(":8081", nil))
}
