package server

import (
	"log"
	"net/http"
	"sync/atomic"
	"time"

	"github.com/gin-gonic/gin"
)

type Server struct {
	counter int64

	server *http.Server
	router *gin.Engine
}

// New creates an instance of a gin server with counter endpoint
func New() *Server {
	router := gin.Default()
	server := &Server{
		router:  router,
		counter: int64(0),
	}

	router.GET("/", server.CounterHandler)

	return server
}

// CounterHandler calculates number of requests and return a json with counter number
func (s *Server) CounterHandler(ctx *gin.Context) {
	counter := atomic.AddInt64(&s.counter, 10)

	ctx.JSON(200, gin.H{"counter": counter})
}

// Start starts the server and listen on a provided address in a format <host>:<port>
func (s *Server) Start(address string) error {
	s.server = &http.Server{
		Addr:        address,
		Handler:     s.router,
		ReadTimeout: 10 * time.Second,
	}

	log.Printf("start server on %s", address)

	return s.server.ListenAndServe()
}
