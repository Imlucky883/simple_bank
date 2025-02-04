package api

import (
	db "github.com/Imlucky883/simple_bank/db/sqlc"
	"github.com/gin-gonic/gin"
)

// Server serves requests for our Banking service
type Server struct {
	store  *db.Store
	router *gin.Engine
}

// NewServer initializes a new Server with the given store and router
func NewServer(store *db.Store) *Server {
	router := gin.Default()
	server := &Server{store: store, router: router}

	// Set up routes
	server.setupRoutes()

	return server
}

// setupRoutes defines all API routes
func (server *Server) setupRoutes() {
	router := server.router

	// Account routes
	router.POST("/accounts", server.createAccount)        // Create account
	router.GET("/accounts/:id", server.getAccount)        // Get account by ID
	router.GET("/accounts", server.listAccounts)          // List accounts with pagination
	router.DELETE("/accounts/:id", server.deleteAccount)  // Delete account by ID
	server.router.POST("/users", server.createUser)       // Create user
	server.router.GET("/users/:username", server.getUser) // Get user by username
}

// Start runs the HTTP server on the given address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
