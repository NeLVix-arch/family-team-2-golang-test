// family team 2 golang test project

package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"family-team/src/config"
	"family-team/src/files"
	"family-team/src/logging"
	"family-team/src/server"

	"go.uber.org/zap"
)

func main() {
	// parse the command line parameters
	c := config.ParseConfig()

	// create logger
	logger := logging.CreateLogger()
	defer logger.Sync()

	// open ZIP file and get the list of files
	files, close := files.OpenZipFile(logger, c.File, c.Ext)
	defer close()

	// create server
	s := server.NewServer(logger, files, c.Port)
	service := s.Build()

	// run server in a goroutine
	go func() {
		// start the HTTP server
		s.Logger.Info("starting HTTP server", zap.String("port", s.Port))
		if err := service.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			s.Logger.Fatal("failed to start HTTP server", zap.Error(err))
		}
	}()

	// handle Ctrl+C signals
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	<-sig
	logger.Info("program is exiting")

	ctx, ctxc := context.WithTimeout(context.Background(), time.Second*10)
	defer ctxc()

	if err := service.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}
