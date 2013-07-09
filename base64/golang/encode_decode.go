package main

import (
    "fmt"
    "encoding/base64"
    "io/ioutil"
    "os"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Please specify filename")
        os.Exit(1)
    }
    file_name := os.Args[1] 
    d1, err := ioutil.ReadFile(file_name)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    str := base64.StdEncoding.EncodeToString(d1)
    d2,err := base64.StdEncoding.DecodeString(str)
    if err != nil {
        fmt.Println(err)
        os.Exit(1) 
    }
    err = ioutil.WriteFile(file_name + ".new", d2, 0755)
    if err != nil {
        fmt.Println(err)
        os.Exit(1) 
    }
    fmt.Println("Done.")
}
