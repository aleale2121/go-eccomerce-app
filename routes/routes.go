package routes

import (
	"embed"
	"io/fs"
	"log"
	"net/http"

	"github.com/aleale2121/go-eccomerce-app/middleware"
	"github.com/gorilla/mux"
)

func NewRouter(staticDir embed.FS, h Routes) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", h.homeGetHandler).Methods("GET")
	r.HandleFunc("/", h.homePostHandler).Methods("POST")
	r.HandleFunc("/register", h.registerGetHandler).Methods("GET")
	r.HandleFunc("/register", h.registerPostHandler).Methods("POST")
	r.HandleFunc("/login", h.loginGetHandler).Methods("GET")
	r.HandleFunc("/login", h.loginPostHandler).Methods("POST")
	r.HandleFunc("/logout", h.logoutGetHandler).Methods("GET")
	r.HandleFunc("/admin", middleware.AuthRequired(h.adminGetHandler)).Methods("GET")
	r.HandleFunc("/products", middleware.AuthRequired(h.productGetHandler)).Methods("GET")
	r.HandleFunc("/product-create", middleware.AuthRequired(h.productCreateGetHandler)).Methods("GET")
	r.HandleFunc("/product-create", middleware.AuthRequired(h.productCreatePostHandler)).Methods("POST")
	r.HandleFunc("/product-edit", middleware.AuthRequired(h.productEditGetHandler)).Methods("GET")
	r.HandleFunc("/product-edit", middleware.AuthRequired(h.productEditPostHandler)).Methods("POST")
	r.HandleFunc("/product-delete", middleware.AuthRequired(h.productDeleteGetHandler)).Methods("GET")
	r.HandleFunc("/users", middleware.AuthRequired(h.userGetHandler)).Methods("GET")
	serverRoot, err := fs.Sub(staticDir, "assets")
	if err != nil {
		log.Fatal(err)
	}
	fileServer := http.FileServer(http.FS(serverRoot))
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServer))
	return r
}
