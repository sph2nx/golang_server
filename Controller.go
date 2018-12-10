package main

import (
    "log"
    "net/http"
    "time"
)

const HostName = "http://694682.ichengyun.net"

var port map[int]string = map[int]string{
    0: "8081",
    1: "8082",
    2: "8083",
}

type RedirectHandler struct{}

func (re *RedirectHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    TimeString := time.Now().Format("2006-01-02 15:04:05.999")
    i := int(TimeString[22]) % 3
    url := HostName + ":" + port[i] + "/"    
    http.Redirect(w, r, url, 307)
}

func main() {
	var re RedirectHandler
	err := http.ListenAndServe(":8080", &re)
	if err != nil {
		log.Fatal(err)
	}
}