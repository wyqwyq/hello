package main

import (
    "fmt"
    "io/ioutil"
    // "log"
    "net/http"
)


func main() {
    resp, err := http.Get("http://o8cpu8afd.bkt.clouddn.com/Vintage.png")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Printf("Wrong here!")
    }else {
        fmt.Println(len(body))
    }
    
}