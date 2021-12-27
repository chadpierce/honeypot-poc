package main

import (
	"os"
	"log"
)

var verNum string = "0.1"
var logName string = "./connections.log"
var detailLogPath string = "./logs"

func main() {

	cfg := Read()

	// TODO log file path is hardcoded to binary path - should this be configurable?
	if cfg.LOGS.Enabled == true {
		
		logfile, err := os.OpenFile(logName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		log.Fatal(err)
	}
		log.SetOutput(logfile)
		defer logfile.Close()  // is this needed here?
	}

	if _, err := os.Stat(detailLogPath); os.IsNotExist(err) {
		os.Mkdir(detailLogPath, 0700)
		log.Println("Created log dir: ", detailLogPath)
	}

	log.Println("### HONEYPOT VERSION", verNum, "STARTED ###")
	tcpServer := NewServer(cfg.TCP.Ports, cfg.LOGS.Detailed)
	tcpServer.Start()
}
