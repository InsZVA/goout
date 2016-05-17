package main

import (
    "../common"
    "net/http"
    "log"
    "flag"
)

var port = flag.String("port", "8888", "Listen port.")

func main() {
    flag.Parse()
    http.ListenAndServeTLS(":" + *port, "../common/cert.pem", "../common/key.pem", &common.Server{})
    log.Println("Server end.")
}