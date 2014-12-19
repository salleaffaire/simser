package main

import ( 
  "fmt"
  "io/ioutil"
  "net/http"
)

func loadPage(filename string) ([]byte, error) {
  body, err := ioutil.ReadFile(filename)
  return body, err
}

func handler(w http.ResponseWriter, r *http.Request) {
  body, err := loadPage(r.URL.Path[1:])
  if err != nil { 
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
  fmt.Fprintf(w, "%s", body)
}

func main() {
  http.HandleFunc("/", handler)
  http.ListenAndServe(":8080", nil)
}
