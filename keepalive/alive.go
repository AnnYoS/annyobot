package keepalive

import (
  "fmt"
  "log"
  "net/http"
)

 func KeepAlive(){

  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
    fmt.Println("Hello, I'm alive")
  })

  log.Fatal(http.ListenAndServe(":8080", nil))
}