package main

import (
	"crypto/rand"
	"fmt"
	"html/template"
	"net/http"
	"time"

	"./api"
	"./db/documents"
	"./models"
	"github.com/gorilla/mux"
	mgo "gopkg.in/mgo.v2"
)

var postsCollection *mgo.Collection

func generateID() string {
	b := make([]byte, 16)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

var templates *template.Template

//IndexHandler dada
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	postDocuments := []documents.PostDocument{}
	postsCollection.Find(nil).All(&postDocuments)
	posts := []models.Post{}
	for _, doc := range postDocuments {
		post := models.Post{doc.ID, doc.Title, doc.Content, doc.Date}
		posts = append(posts, post)
	}
	RenderTemplate(w, "./views/index.html", posts)
}

//AdminHandler dada
func adminHandlerPost(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	content := r.FormValue("content")
	time := time.Now().UTC()
	id := r.FormValue("id")

	postDocument := documents.PostDocument{id, title, content, time}
	if id != "" {
		postsCollection.UpdateId(id, postDocument)

	} else {
		id = generateID()
		postDocument.ID = id
		postsCollection.Insert(postDocument)
	}

	http.Redirect(w, r, "/", 302)
}

func adminHandlerGet(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "./views/admin.html", nil)
}

//RenderTemplate Template rendering function
func RenderTemplate(w http.ResponseWriter, templateFile string, templateData interface{}) {
	t, err := template.ParseFiles(templateFile, "./views/partials/head.html", "./views/partials/footer.html")
	if err != nil {
		fmt.Fprintln(w, "not implemented yet !", err)
	}
	t.Execute(w, templateData)
}

func main() {

	session, err := mgo.Dial("localhost")
	defer session.Close()
	if err != nil {
		panic(err)
	}
	postsCollection = session.DB("movies").C("movies")
	templates = template.Must(template.ParseGlob("views/*.html"))

	r := mux.NewRouter()
	r.HandleFunc("/", IndexHandler).Methods("GET")
	r.HandleFunc("/admin", adminHandlerPost).Methods("POST")
	r.HandleFunc("/admin", adminHandlerGet).Methods("GET")
	r.HandleFunc("/movies", api.AllMoviesEndPoint).Methods("GET")
	r.HandleFunc("/movies", api.CreateMovieEndPoint).Methods("POST")
	r.HandleFunc("/movies", api.UpdateMovieEndPoint).Methods("PUT")
	r.HandleFunc("/movies", api.DeleteMovieEndPoint).Methods("DELETE")
	r.HandleFunc("/movies/{id}", api.FindMovieEndpoint).Methods("GET")

	fs := http.FileServer(http.Dir("./static/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
