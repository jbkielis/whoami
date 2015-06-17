package main

import (
  "os"
  "fmt"
  "net/http"
  "log"
)

func handler(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Server", "A Go Web Server")
  w.WriteHeader(200)  
  w.Write([]byte("OK"))
}

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    http.HandleFunc("/health", handler)

    fmt.Fprintf(os.Stdout, "Listening on :%s\n", port)
    hostname, _ := os.Hostname()
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(os.Stdout, "I'm %s\n", hostname)
    	fmt.Fprintf(w, "I'm %s\n", hostname)
    })


    log.Fatal(http.ListenAndServe(":" + port, nil))
}

