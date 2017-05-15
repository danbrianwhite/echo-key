package main

import (
	"flag"
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "%s", *strFlag)
	fmt.Printf("Key Sent\n")
}

var strFlag = flag.String("key", "missing SSH key", "SSH Key")

func main() {
	port := flag.String("port", "8080", "port")
	flag.Parse()
	
	fmt.Printf("echo-key running")

	portString := *port
    http.HandleFunc("/", handler)
    http.ListenAndServe(":" + portString, nil)
}