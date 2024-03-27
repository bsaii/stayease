package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/bsaii/stayease/internal/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	_ "ariga.io/atlas-provider-gorm/gormschema"
	_ "github.com/bsaii/stayease/docs"
)

func SetDBMiddleware(db *gorm.DB) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			timeoutContext, cancel := context.WithTimeout(context.Background(), 30*time.Second)
			defer cancel()
			ctx := context.WithValue(r.Context(), "db", db.WithContext(timeoutContext))
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

//	@title			StayEase API
//	@version		1.0
//	@description	StayEase is a comprehensive room management API that simplifies the process of booking and managing rooms. With StayEase, users can easily search for available rooms, make reservations, and perform various management tasks related to room bookings.
//	@termsOfService	http://stayease.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.stayease.io/support
//	@contact.email	support@stayease.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:8080
//	@BasePath	/

//	@securityDefinitions.basic	BasicAuth

// @externalDocs.description	OpenAPI
// @externalDocs.url			https://swagger.io/resources/open-api/
func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	dbConnString := os.Getenv("DB_CONNECTION_STRING")
	if dbConnString == "" {
		log.Fatal("DB_CONNECTION_STRING environment variable not set.")
	}

	db, err := gorm.Open(postgres.Open(dbConnString), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(SetDBMiddleware(db))

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to StayEase!"))
	})
	r.Route("/rooms", func(r chi.Router) {
		r.Get("/", handler.AllRooms)
		r.Post("/", handler.AddRoom)

		r.Route("/{roomID}", func(r chi.Router) {
			r.Get("/", handler.GetRoom)
			r.Put("/", handler.UpdateRoom)
			r.Delete("/", handler.DelRoom)
		})
	})
	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"),
	))

	fmt.Println("Starting server on port:8080...")
	http.ListenAndServe(":8080", r)
}
