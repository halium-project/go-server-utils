package server

import (
	"log"
	"net/http"
	"time"

	"github.com/phyber/negroni-gzip/gzip"
	"github.com/rs/cors"
	"github.com/urfave/negroni"
	"gopkg.in/tylerb/graceful.v1"
)

func ServeHandler(addr string, handler http.Handler) {
	server := negroni.New().
		With(negroni.NewRecovery()).
		With(negroni.NewLogger()).
		With(gzip.Gzip(gzip.DefaultCompression)).
		With(cors.Default())

	server.UseHandler(handler)

	log.Printf("start listening on %s", addr)
	graceful.Run(addr, 10*time.Second, server)
}