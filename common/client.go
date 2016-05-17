package common

import (
    "net/http"
    "crypto/tls"
    "net"
    "io"
    "bufio"
)

type Client struct {
    remote string
    user   string
    passwd string
}

func NewClient(remote string, user string, passwd string) *Client {
    return &Client {
        remote: remote,
        user: user,
        passwd: passwd,
    }
}

func (c *Client) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    conn, err := net.Dial("tcp", c.remote)
    defer conn.Close()
    
    if err != nil {
        return
    }
    conns := tls.Client(conn, &tls.Config{InsecureSkipVerify: true})
    r.WriteProxy(conns)
    
    reader := bufio.NewReader(conns)
    resp, err := http.ReadResponse(reader, r)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	for k, v := range resp.Header {
        for _, vv := range v {
            w.Header().Add(k, vv)
        }
    }
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}