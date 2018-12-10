package main

import (
	"fmt"
	"log"
	"net/http"
)

type PrintConn struct{}

func (p *PrintConn) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Connected to port: 8081!")
}

func main() {
	var p PrintConn
	err := http.ListenAndServe(":8081", &p)
	if err != nil {
		log.Fatal(err)
	}
}
