package main

import (
	"fmt"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/none-da/go_contexts/pkg/server"
)

const port = 8080

func init() {
	log.SetLevel(log.DebugLevel)
	log.SetFormatter(&log.TextFormatter{
		DisableLevelTruncation: true,
		PadLevelText:           true,
		TimestampFormat:        time.RFC3339,
		FullTimestamp:          true,
		ForceColors:            false,
	})
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Debug("request started")
	fmt.Fprintf(w, server.Hello())
	log.Debug("request finished")
}

func main() {
	http.HandleFunc("/", handler)
	log.Info("Server started listening on ", port)
	http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
}
