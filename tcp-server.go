package main

import (
	//"fmt"
	"log"
	"net"
	"sync"
	"encoding/base64"
	"strings"
	"os"
)

// Server is the tcp server struct
type Server struct {
	Ports []string
	DetailedLogging bool
}

func encode64(decodedString string) string {
	// if isBase64 == false {
	// 	return decodedString
	// }
	var encodedString = base64.StdEncoding.EncodeToString([]byte(decodedString))
	return encodedString
}

func decode64(encodedString string) string {
	// if isBase64 == false {
	// 	return encodedString
	// }
	var decodedByte, _ = base64.StdEncoding.DecodeString(encodedString)
	var decodedString = string(decodedByte)
	return decodedString
}

func getIPLogFileName(remoteAddr string) string {
	if remoteAddr == "::1" {
 		return "localhost.txt"
	}
	uip := strings.Replace(remoteAddr, ".", "_", -1)
	fname := uip + ".txt"
	return fname
  }

func NewServer(ports []string, detailedLogging bool) *Server {
	return &Server{ports, detailedLogging}
}

func (t *Server) Start() {
	var wg sync.WaitGroup
	wg.Add(len(t.Ports))
	for _, port := range t.Ports {
		go func(port string, wg *sync.WaitGroup) {
			log.Println("TCP SERVER STARTED ON PORT:", port)
			listen, err := net.Listen("tcp", ":"+port)
			if err != nil {
				log.Println("ERROR port listener: ", err)
				wg.Done()
				return
			}
			for {
				conn, err := listen.Accept()
				if err != nil {
					log.Fatal(err)
				}
				go handleConnection(conn, t.DetailedLogging)
			}
		}(port, &wg)
	}
	wg.Wait()
	log.Println("TCP SERVER STOPPED")
}

func handleConnection(conn net.Conn, detail bool) {
	data := make([]byte, 4096)
	n, err := conn.Read(data)
	if err != nil {
		log.Println("ERROR handle conn", err)
		conn.Close()
		return
	}
	defer conn.Close()
	
	RHost, RPort, err := net.SplitHostPort(conn.RemoteAddr().String())
	if err != nil {
		log.Println("ERROR parse remote host:", err)
		return
	}
	LHost, LPort, err := net.SplitHostPort(conn.LocalAddr().String())
	if err != nil {
		log.Println("ERROR parse local host:", err)
		return
	}
	// TODO fix log output
	// TODO how to format detailed output data
	// TODO path should be taken from config
	// FIXME why does data have trailing A's without using data[:n]
	log.Println("CONNECTION:", RHost, RPort, LHost, LPort, n, encode64(string(data[:n])))
	if detail == true {
		detailedOutput := "DETAILED STUFF GOES HERE\n"
		//log.Println("DETAILED LOGGING TRUE")
		fname := getIPLogFileName(RHost)
		f, err := os.OpenFile(fname, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			log.Println("ERROR opening or creating detailed logfile for", RHost, err)
		} else {
			if _, err = f.WriteString(detailedOutput); err != nil {
				log.Println("ERROR writing to logfile for", RHost, err)
			} else {
				log.Println("Detailed log written for", RHost)
			}
		}
		defer f.Close()
	}
}
