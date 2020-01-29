package main

import(
  "fmt"
  "log"
  "os"
  "net/http"
  "github.com/aleale2121/go-webapp/routes"
  "github.com/aleale2121/go-webapp/utils"
  "github.com/aleale2121/go-webapp/models"
  "github.com/aleale2121/go-webapp/sessions"
)


func main() {
  models.TestConnection()
  sessions.SessionOptions("localhost", "/", 3600, true)
  port := os.Getenv("PORT")
  if port == "" {
    port = "8080"
  }
  fmt.Printf("Listening Port %s\n", port)
  utils.LoadTemplates("views/*.html")
  r := routes.NewRouter()
  http.Handle("/", r)
  log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

