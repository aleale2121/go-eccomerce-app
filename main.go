package main

import (
	"embed"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/aleale2121/go-eccomerce-app/models"
	"github.com/aleale2121/go-eccomerce-app/routes"
	"github.com/aleale2121/go-eccomerce-app/sessions"
	"github.com/aleale2121/go-eccomerce-app/utils"
)

//go:embed views/*
//go:embed assets/*
var content embed.FS

func main() {
	models.TestConnection()
	sessions.SessionOptions("localhost", "/", 3600, true)

	port := os.Getenv("PORT")
	if port == "" {
		fmt.Println("Not Port")
		os.Exit(1)
	}
	models.CreateCat()
	models.CreatePro()
	models.CreateUser()
	fmt.Printf("Listening Port %s\n", port)
	utils.LoadTemplates("views/*.html")
	r := routes.NewRouter(content)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
