package server

import (
	dcodec "urlshortener/codec"
	"urlshortener/storage"
)

// RedirectServer - URL shortener HTTP server structure with bound methods
type RedirectServer struct {
	codec       *dcodec.URLCodec
	storage     storage.GenericStorage
	templateDir string
}
