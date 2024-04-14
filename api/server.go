package api

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	db "github.com/naviscom/catalystx2/db/sqlc"
	"github.com/naviscom/catalystx2/token"
	"github.com/naviscom/catalystx2/util"
)

// // Server serves HTTP requests.
type Server struct {
	config     util.Config
	store      *db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

// NewServer creates a new HTTP Gin server and set up routing & CORS.
func NewServer(config util.Config, store *db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}
	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()
	//router.POST("/users", server.createUser)
	//router.POST("/users/login", server.loginUser)
	//router.POST("/tokens/renew_access", server.renewAccessToken)
	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))

	// setting the Gin router to Default() allows all origins
	router.Use(cors.Default())

	server.router = router
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
