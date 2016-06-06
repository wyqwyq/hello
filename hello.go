package main

import "fmt"

func main() {
    s := [] int {1, 2, 3}
    arr := s[:len(s) - 1]
    if &s[0] == &arr[0]{
        fmt.Printf("Only Ref\n")
    }else{
        fmt.Printf("Copy value")
    }
    fmt.Printf("hello, world\n")
}