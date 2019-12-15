package server

import (
	dcodec "urlshortener/codec"
	"urlshortener/storage"
)

type RedirectServer struct {
	codec       *dcodec.URLCodec
	storage     storage.GenericStorage
	templateDir string
}
