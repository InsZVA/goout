package common

import (
    "net/http"
    "io/ioutil"
    "time"
)

var transport = &http.Transport{
	ResponseHeaderTimeout: 30 * time.Second,
}

type Server struct{
    code int
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    resp, err := transport.RoundTrip(r)
    for k, v := range resp.Header {
        for _, vv := range v {
            w.Header().Add(k, vv)
        }
    }
    if err != nil {
        panic(err)
    }
    data, _ := ioutil.ReadAll(resp.Body)
    defer resp.Body.Close()
    w.WriteHeader(resp.StatusCode)
    w.Write(data)
}