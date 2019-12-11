package server

import (
	dcodec "urlshortener/codec"
	"urlshortener/urlstorage"
)

type RedirectServer struct {
	codec       *dcodec.URLCodec
	storage     *urlstorage.URLStorage
	templateDir string
}
