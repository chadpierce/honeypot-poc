package main

import (
	"os"
	"log"
)

func main() {

	cfg := Read()

	if cfg.LOG.Enabled == true {
		// TODO fix log path in congig to be ONLY path. log file is separate item or hard coded
		// use same path for detailed log
		logfile, err := os.OpenFile(cfg.LOG.Path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
		log.SetOutput(logfile)
		defer logfile.Close()  // is this needed here?
	}

	tcpServer := NewServer(cfg.TCP.Ports, cfg.LOG.Detailed)
	tcpServer.Start()
}
