package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/rs/zerolog"

	"github.com/a8filich/graph-collab-api/internal/restful"
)

var (
	logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}).
		Level(zerolog.DebugLevel).
		With().
		Timestamp().
		Caller().
		Logger().
		Sample(&zerolog.BasicSampler{N: 10})

	host = "127.0.0.1"
	port = 8080
)

func main() {
	router := restful.Router()

	addr := fmt.Sprintf("%s:%d", host, port)

	logger.Debug().Msgf("Starting to run on %s", addr)

	if err := http.ListenAndServe(addr, router); err != nil {
		logger.Panic().Msg(err.Error())
	}
}
