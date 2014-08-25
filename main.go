package main

import (
  "fmt"
  "net/http"
  "github.com/pilu/traffic"
)

func rootHandler(w traffic.ResponseWriter, r *traffic.Request) {
  w.Render("index", w.GetVar("phone"))
}

func serverSideHandler(w traffic.ResponseWriter, r *traffic.Request) {
  w.WriteText("Resource not available\n")
}

func addLocationHeader(w traffic.ResponseWriter, r *traffic.Request) {
  t := fmt.Sprintf("tel://%s", w.GetVar("phone"))
  w.Header().Add("Location",  t)
  w.WriteHeader(http.StatusTemporaryRedirect)
}

func customNotFoundHandler(w traffic.ResponseWriter, r *traffic.Request) {
  w.WriteHeader(http.StatusNotFound)
  w.WriteText("Page not found: %s\n", r.URL.Path)
}

func main() {
  router := traffic.New()
  router.Get("/", rootHandler)
  router.Get("/server_side/", serverSideHandler).AddBeforeFilter(addLocationHeader)
  router.NotFoundHandler = customNotFoundHandler
  router.Run()
}
