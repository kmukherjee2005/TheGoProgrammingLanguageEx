package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"path/filepath"

	"github.com/kmukherjee2005/TheGoProgrammingLanguageEx/Ch08/ex02/ftp"
)

var port int
var rootDir string

func init() {
	flag.IntVar(&port, "Port", 8080, "Port Number")
	flag.StringVar(&rootDir, "rootDir", ".", "Root Dir")

	flag.Parse()
}

func main() {
	server := fmt.Sprintf("%d", port)
	listener, err := net.Listen("tcp", server)

	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()

		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}
func handleConn(c net.Conn) {
	defer c.Close()

	abs, err := filepath.Abs(rootDir)

	if err != nil {
		log.Fatal(err)
	}
	ftp.Serve(ftp.NewConn(c, abs))

}
