package server

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jffhtchr/record_project/cmd/server/middleware"
	"github.com/jffhtchr/record_project/pkg/controllers"
)

func NewRouter(db *sql.DB) *chi.Mux {
	r := chi.NewRouter()
	registerRoutes(r, db)
	return r
}

func StartServer(db *sql.DB) {
	r := NewRouter(db)
	log.Println(fmt.Printf("Starting Server on port %s", "8000"))
	log.Fatal(http.ListenAndServe(":8000", middleware.JsonContentTypeMiddleware(r)))
}

func registerRoutes(router chi.Router, db *sql.DB) {
	router.Route("/", func(r chi.Router) {
		r.Get("/ping", controllers.Ping)               // GET  health
		r.Get("/albums", controllers.GetAllAlbums(db)) // GET  albums
		r.Post("/album", controllers.CreateAlbum(db))  // POST album

	})
}
