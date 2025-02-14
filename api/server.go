package api

import (
	"fmt"

	db "github.com/Imlucky883/simple_bank/db/sqlc"
	"github.com/Imlucky883/simple_bank/db/util"
	"github.com/Imlucky883/simple_bank/token"
	"github.com/gin-gonic/gin"
)

// Server serves requests for our Banking service
type Server struct {
	config     util.Config // Configuration for the server and database connection
	store      *db.Store
	router     *gin.Engine
	tokenMaker token.Maker
}

// NewServer initializes a new Server with the given store and router
func NewServer(config util.Config, store *db.Store) *Server {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		panic(fmt.Errorf("cannot create token maker: %w", err))
	}

	router := gin.Default()
	server := &Server{config: config, store: store, router: router, tokenMaker: tokenMaker}

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
	server.router.POST("/users/login", server.loginUser)  // Login
}

// createToken handles the creation of a new token
func (server *Server) createToken(ctx *gin.Context) {
	// Implementation for creating a token
}

// Start runs the HTTP server on the given address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
