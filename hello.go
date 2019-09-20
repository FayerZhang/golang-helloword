package main

import (
    "io"
    "net/http"
    "log"
    "os"
)

// hello world, the web server
func HelloServer(w http.ResponseWriter, req *http.Request) {
    version := os.Getenv("version")
    SYS_INCREMENT := os.Getenv("SYS_INCREMENT")
    SYS_TIMESTAMP := os.Getenv("SYS_TIMESTAMP")
    io.WriteString(w, "hello, world!\n")
    io.WriteString(w, "version="+ version + "\n")
    io.WriteString(w, "SYS_INCREMENT="+ SYS_INCREMENT + "\n")
    io.WriteString(w, "SYS_TIMESTAMP="+ SYS_TIMESTAMP + "\n")
}

func main() {
    http.HandleFunc("/hello", HelloServer)
    err := http.ListenAndServe(":12345", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
