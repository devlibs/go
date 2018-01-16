package main 

import (
    "encoding/json"
    "fmt"
)

func main() {
    str := `{"changes": [{"armid":3,"Index":5}, {"armid":3,"Index":6}]}`
    var msg map[string]interface{}
    err := json.Unmarshal([]byte(str), &msg)
    if err!=nil {
        fmt.Println("fuck",err)
    }
    fmt.Println(msg)
    
    changes,ok := msg["changes"].([](interface{}))
    if !ok {
        fmt.Println("Can't convert msg to []interface")
    }
    
    for _ ,ichange := range changes {
        change, ok := ichange.(map[string]interface{})
          if !ok {
              fmt.Println("cant convert ichange to map[string]interface{}")
          }
        Index, ok := change["Index"].(float64)
        if !ok {    
            fmt.Println("cant convert Index to float64")
        }
        armid,ok := change["armid"].(float64)
        if !ok {    
            fmt.Println("cant convert armid to float")
        }
        fmt.Println(Index,armid)
    }
}