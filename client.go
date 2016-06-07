package main

import (
    "fmt"
    "os"
    // "io/ioutil"
    // "log"
    "net/http"
)



func main() {
    file_path := http.Dir(".")
    file, _ := file_path.Open("client.go")
    buf := [1024]byte{}
    fmt.Printf(string(buf))
}