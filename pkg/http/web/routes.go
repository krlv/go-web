package web

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/krlv/goweb/pkg/storage"
)

// DB is a page repository
var db *storage.MemoryStorage

func init() {
	db = storage.NewStorage()
}

// NotFound will be invoked by mux when URL can't be matched
// to registered routes (404 page)
func NotFound(w http.ResponseWriter, r *http.Request) {
	log.Print("NotFoundHandler:", r.RequestURI)
	path := strings.TrimLeft(r.RequestURI, "/")

	t, _ := template.ParseFiles("templates/404.html")
	t.Execute(w, map[string]interface{}{"Path": path})
}

// StaticPage renders static pages
func StaticPage(w http.ResponseWriter, r *http.Request) {
	path, _ := mux.CurrentRoute(r).GetPathTemplate()
	path = strings.TrimLeft(path, "/")

	if path == "" {
		path = "home"
	}

	fileName := "public/" + path + ".html"

	http.ServeFile(w, r, fileName)
}

// GetPage renders list of pages from data storage
func GetPages(w http.ResponseWriter, r *http.Request) {
	pages := db.FindPages()

	t, err := template.ParseFiles("templates/blog.html")
	if err != nil {
		// TODO handle
	}

	t.Execute(w, pages)
}

// GetPageBySlug returns a single page by it's slug
func GetPageBySlug(w http.ResponseWriter, r *http.Request) {
	slug := mux.Vars(r)["slug"]

	p, err := db.GetPageBySlug(slug)
	if err != nil {
		// TODO handle
	}

	t, err := template.ParseFiles("templates/post.html")
	if err != nil {
		// TODO handle
	}

	t.Execute(w, p)
}

// GetNotes returns list of notes from data storage
func GetNotes(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/notes.html")
	if err != nil {
		// TODO handle
	}

	err = t.Execute(w, db.FindNotes())
	if err != nil {
		// TODO handle
	}
}

// GetNote returns a single note by note ID
func GetNote(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		// TODO handle
	}

	n, err := db.GetNoteByID(id)
	if err != nil {
		// TODO handle
	}

	t, err := template.ParseFiles("templates/note.html")
	if err != nil {
		// TODO handle
	}

	t.Execute(w, n)
}
