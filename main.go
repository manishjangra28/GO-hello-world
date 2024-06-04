package main

import "net/http"

func main() {
 http.HandleFunc("/hello", helloword)

 http.ListenAndServe(":9000", nil)
}

func helloword(w http.ResponseWriter, r *http.Request) {
 w.Write([]byte("Hello from Botgauge"))
}