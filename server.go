package main

import (
    "io"
    "net/http"
)

func hello(rw http.ResponseWriter, req *http.Request) {
    io.WriteString(rw, "hello widuu")
}

func main() {
    http.HandleFunc("/", hello)  //设定访问的路径
    http.ListenAndServe(":8090", nil) //设定端口和handler
}
