package server

import (
	"fmt"
	"net"
	"os"

	"github.com/sandeepravi/muninndb/database"
)

var (
	s *Server
)

// Server is the basic db server object
type Server struct {
	l net.Listener
	//cmds    map[string]*command
	dbs     map[int]*database.DB
	clients map[*Client]bool
	cmds    map[string]*Command
	cfg     *Config
}

func Setup() {
	s = &Server{
		cmds:    make(map[string]*Command),
		dbs:     make(map[int]*database.DB),
		clients: make(map[*Client]bool),
	}

	s.cfg = &Config{
		host: "localhost",
		port: "7379",
	}
}

// Start runs the server and sets up config
func Start() {
	fmt.Println("Starting server on " + s.cfg.host + ":" + s.cfg.port + " ...")
	// Listen for tcp connections on port
	l, err := net.Listen("tcp", s.cfg.host+":"+s.cfg.port)
	if err != nil {
		fmt.Println("Error listening: ", err.Error())
		os.Exit(1)
	}

	// Close listener when connection closes
	defer l.Close()

	for {
		c, err := l.Accept()

		if err != nil {
			fmt.Println("Error acception conenction: ", err.Error())
			os.Exit(1)
		}

		Handler(c, s)
		// Close connection
		c.Close()
	}
}
