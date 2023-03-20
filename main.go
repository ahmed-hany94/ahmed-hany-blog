package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

// *****************
// Date Structures *
// *****************

type Post struct {
	Title   string
	Slug    string
	Date    string
	Tags    string
	Content template.HTML
}

type Posts struct {
	PostsList []Post
}

// *********
// Globals *
// *********

var PostsToList Posts

var templates = template.Must(
	template.ParseFiles(
		"public/index.html",
		"public/post.html",
	),
)
var validPath = regexp.MustCompile("^/(view)/([a-zA-Z0-9-.\u0621-\u064A]+)$")

// ****************
// Util Functions *
// ****************

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

func unslugify(title *string) {
	t := strings.Split(*title, "-")
	*title = strings.Join(t, " ")
}

// ********************
// Handlers Functions *
// ********************

// this will show all posts at /
func listHandler(w http.ResponseWriter, r *http.Request) {
	postsDirectoryList, err := os.ReadDir("posts/")
	check(err)

	if len(PostsToList.PostsList) == 0 {
		for i := len(postsDirectoryList) - 1; i >= 0; i-- {
			p := postsDirectoryList[i]

			title := strings.Replace(p.Name(), ".html", "", 1)

			unslugify(&title)

			PostsToList.PostsList = append(PostsToList.PostsList, Post{Title: title, Slug: p.Name()})
		}

	}

	templates.ExecuteTemplate(w, "index.html", PostsToList)
}

// this will show post at /view/<slug>
func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	path := "./posts/" + title
	post_body, err := os.ReadFile(path)
	check(err)

	post_content := template.HTML(post_body)

	post := Post{
		Content: post_content,
	}

	templates.ExecuteTemplate(w, "post.html", post)
}

// this will show /cv
func viewCVHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./public/cv.html")
}

// this will show /projects
func viewProjectsHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./public/projects.html")
}

// this will show /robots.txt
func viewRobotsHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./public/robots.txt")
}

// this will show /feed.xml
func viewFeedHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./public/feed.xml")
}

// ***************
// Main Function *
// ***************

func main() {
	http.HandleFunc("/", listHandler)
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/cv", viewCVHandler)
	http.HandleFunc("/projects", viewProjectsHandler)
	http.HandleFunc("/feed.xml", viewFeedHandler)
	http.HandleFunc("/robots.txt", viewRobotsHandler)

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	port := os.Getenv("PORT")

	if port == "" {
		port = "4444"
	}

	log.Println("listen on", ":"+port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
