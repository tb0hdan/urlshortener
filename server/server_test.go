package server

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	dcodec "urlshortener/codec"
	"urlshortener/urlstorage"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestRedirectServer_IndexHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/index.html", nil)
	if err != nil {
		t.Fatal(err)
	}

	s := &RedirectServer{templateDir: "../templates"}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(s.IndexHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestRedirectServer_CatchAllHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	s := &RedirectServer{templateDir: "../templates", codec: dcodec.New(),
		storage: &urlstorage.URLStorage{
			URLs:     make([]*urlstorage.URLItem, 0),
			URLsHash: make(map[string]string),
		},
	}
	rr := httptest.NewRecorder()

	router := mux.NewRouter()

	router.ServeHTTP(rr, req)
	router.PathPrefix("/").HandlerFunc(s.CatchAllHandler)
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNotFound)
	}
}

func TestRedirectServer_ShortenHandler(t *testing.T) {
	myURL := "http://example.com"
	shortURL := "cab"
	data := url.Values{"url": {myURL}}

	req, err := http.NewRequest("POST", "/shorten", strings.NewReader(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	if err != nil {
		t.Fatal(err)
	}

	s := &RedirectServer{templateDir: "../templates", codec: dcodec.New(),
		storage: &urlstorage.URLStorage{
			URLs:     make([]*urlstorage.URLItem, 0),
			URLsHash: make(map[string]string),
		},
	}

	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/shorten", s.ShortenHandler)
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	tmpl := template.Must(template.ParseFiles(fmt.Sprintf("%s/shorten.html", s.templateDir)))
	wr := bytes.NewBufferString("")

	_ = tmpl.Execute(wr, fmt.Sprintf("%s -> %s", myURL, shortURL))

	myAssert := assert.New(t)

	myAssert.Equal(wr.String(), rr.Body.String())
}
