package main

import (
	"flag"
    "fmt"
    "log"
    "net"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(200)
    fmt.Fprintf(w, "%s", *strFlag)
	fmt.Printf("Key Sent\n")
}

var strFlag = flag.String("key", "missing SSH key", "SSH Key")

func main() {
	port := flag.String("port", "8080", "port")
	flag.Parse()
	
	fmt.Printf("echo-key running")

	portString := *port

    smHttp := http.NewServeMux()
    smHttp.HandleFunc("/", handler)

    l, err := net.Listen("tcp4", "0.0.0.0:" + portString)
    if err != nil {
        log.Fatal(err)
    }
    log.Fatal(http.Serve(l, smHttp))
}