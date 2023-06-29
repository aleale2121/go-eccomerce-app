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
	"github.com/spf13/viper"

	_ "github.com/lib/pq"
)

//go:embed views/*
//go:embed assets/*
var content embed.FS

// Config stores all configuration of the application.
// The values are read by viper from a config file or environment variable.
type Config struct {
	PORT     string `mapstructure:"PORT"`
	DBSource string `mapstructure:"DATABASE_URL"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig() (config Config, err error) {
	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}

func main() {
	cfg, err := LoadConfig()
	if err != nil {
		fmt.Println("Unable to load config")
		os.Exit(1)
	}
	fmt.Println(cfg)
	db, err := models.Connect(cfg.DBSource)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	fmt.Println("Database Connected Successfully")

	sessions.SessionOptions("localhost", "/", 3600, true)

	if cfg.PORT == "" {
		fmt.Println("Not Port")
		os.Exit(1)
	}

	store := models.NewStore(db)
	store.CreateCat()
	store.CreatePro()
	store.CreateUser()
	fmt.Printf("Listening Port %s\n", cfg.PORT)
	utils.LoadTemplates("views/*.html")
	route := routes.NewRoutes(store)
	r := routes.NewRouter(content, route)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", cfg.PORT), nil))
}
