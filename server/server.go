package server

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
	"time"

	dcodec "urlshortener/codec"
	"urlshortener/miscellaneous"
	"urlshortener/urlstorage"

	"github.com/gorilla/mux"
)

func (rs *RedirectServer) IndexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(fmt.Sprintf("%s/index.html", rs.templateDir)))

	w.Header().Add("Content-Type", "text/html")
	_ = tmpl.Execute(w, "")
}

func (rs *RedirectServer) CatchAllHandler(w http.ResponseWriter, r *http.Request) {
	if r.RequestURI == "/" {
		rs.IndexHandler(w, r)
		return
	}

	shortenedURI := strings.TrimPrefix(r.RequestURI, "/")
	urlID := rs.codec.Encode(shortenedURI) - dcodec.StartValue
	fmt.Println(shortenedURI, urlID)

	url := rs.storage.GetByID(urlID)
	if len(url) == 0 {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "404 Not Found")

		return
	}

	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func (rs *RedirectServer) ShortenHandler(w http.ResponseWriter, r *http.Request) {
	var (
		urlID    int64
		shortURL string
	)

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprint(w, "405 Not Allowed")
	}

	tmpl := template.Must(template.ParseFiles(fmt.Sprintf("%s/shorten.html", rs.templateDir)))

	w.Header().Add("Content-Type", "text/html")

	url := r.FormValue("url")

	// Get last id
	lastID := rs.storage.LenSafe()

	// Check for URL in DB
	if value, ok := rs.storage.GetByLongURL(url); ok {
		// URL in DB
		shortURL = value
	} else {
		// No URL in DB
		urlID = lastID + 1
		shortURL = rs.codec.Decode(urlID + dcodec.StartValue)
		added := rs.storage.Add(shortURL, url)
		if added != 0 {
			fmt.Printf("New URL %s -> %s, #%d", url, shortURL, added)
		} else {
			fmt.Printf("URL %s not added", url)
		}
	}

	_ = tmpl.Execute(w, fmt.Sprintf("%s -> %s", url, shortURL))
}

func Run(serverConfig *miscellaneous.ServerConfig) {
	rs := &RedirectServer{codec: dcodec.New(),
		storage: &urlstorage.URLStorage{
			URLs:     make([]*urlstorage.URLItem, 0),
			URLsHash: make(map[string]string),
		},
	}
	r := mux.NewRouter()
	r.HandleFunc("/index.{string}", rs.IndexHandler)
	r.HandleFunc("/shorten", rs.ShortenHandler)

	r.PathPrefix("/").HandlerFunc(rs.CatchAllHandler)

	srv := &http.Server{
		Handler: r,
		Addr:    serverConfig.Bind,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: time.Duration(int64(serverConfig.WriteTimeout) * time.Second.Nanoseconds()),
		ReadTimeout:  time.Duration(int64(serverConfig.ReadTimeout) * time.Second.Nanoseconds()),
	}

	log.Fatal(srv.ListenAndServe())
}
