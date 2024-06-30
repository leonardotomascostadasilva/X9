package server

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/leonardotomascostadasilva/X9/consumers"
	"github.com/leonardotomascostadasilva/X9/internal/config"
)

type Server struct {
	port int
}

func Init() {

	go consumers.TroubleShootingConsumerExecute()

	port, _ := strconv.Atoi(config.Get().HttpPort)

	NewServer := &Server{
		port: port,
	}

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
