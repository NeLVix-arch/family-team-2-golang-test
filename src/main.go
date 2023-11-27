// family team 2 golang test project

package main

import (
	"os"
	"os/signal"

	"family-team/src/config"
	"family-team/src/files"
	"family-team/src/logging"
	"family-team/src/server"
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

	// run server in a separate goroutine
	go s.Run()

	// handle Ctrl+C signals
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	<-sig
	logger.Info("program is exiting")
}
