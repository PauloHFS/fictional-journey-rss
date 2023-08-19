package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	var port string = os.Getenv("PORT")

	if port == "" {
		log.Fatal("Port is undefined on .env")
	}

	router := chi.NewRouter()
	corsHandler := cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	})

	router.Use(middleware.Logger)
	router.Use(corsHandler)

	v1Router := chi.NewRouter()
	v1Router.Get("/", (func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	}))

	router.Mount("/v1", v1Router)

	log.Printf("Serving on port: %s\n", port)
	http.ListenAndServe(":"+port, router)
}
