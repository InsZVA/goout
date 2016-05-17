package main

import (
    "../common"
    "net/http"
    "log"
    "flag"
)

var port = flag.String("port", "1080", "Listen port.")
var remote = flag.String("remoteAddr", "127.0.0.1:8888", "Remote server address.")

func main() {
    flag.Parse()
    http.ListenAndServe(":" + *port, common.NewClient(*remote, "InsZVA", "123"))
    log.Println("Server started.")
}