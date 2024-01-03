package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	// Internal Modules

	"github.com/ErayCep/ryg/internal/handlers"
	"github.com/ErayCep/ryg/internal/storage"

	// Third Party Modules
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db_name := os.Getenv("DB_NAME")
	db_user := os.Getenv("DB_USER")
	db_pass := os.Getenv("DB_PASSWORD")

	sql_source := "postgres://" + db_user + ":" + db_pass + "@localhost/" + db_name

	db, err := sql.Open("postgres", sql_source)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	l := log.New(os.Stdout, "ryg", log.LstdFlags)

	strg := storage.Storage{
		DB: db,
	}

	router := mux.NewRouter()

	handler := handlers.NewHandler(l, strg)
	getRouter := router.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/games", handler.GetGamesHandler)
	getRouter.HandleFunc("/games/{id:[0-9]+}", handler.GetGameHandler)
	getRouter.HandleFunc("/reviews", handler.GetReviewsHandler)
	getRouter.HandleFunc("/reviews/{id:[0-9]+}", handler.GetReviewHandler)

	postRouter := router.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/games", handler.PostGamesHandler)

	putRouter := router.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/games/{id:[0-9]+}", handler.PutGamesHandler)

	deleteRouter := router.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc("/games/{id:[0-9]+}", handler.DeleteGamesHandler)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
