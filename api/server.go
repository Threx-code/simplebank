package api

import (
	"fmt"

	db "github.com/Threx-code/simple_bank/db/sqlc"
	"github.com/Threx-code/simple_bank/token"
	"github.com/Threx-code/simple_bank/utils"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type Server struct {
	config     utils.Config
	store      *db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

//NewJWTMaker

func NewServer(config utils.Config, store *db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create to maker %v:", err)
	}

	// registering a custom validator
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	server.setUpRouter()
	return server, nil
}

func (server *Server) setUpRouter() {
	router := gin.Default()

	router.POST("/users", server.createUser)
	router.POST("/users/login", server.loginUser)

	// adding authenticated middleware

	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))
	authRoutes.POST("/accounts", server.createAccount)
	authRoutes.GET("/accounts/:id", server.getAccount)
	authRoutes.GET("/accounts", server.listAccount)
	authRoutes.POST("/users/transfers", server.createTransfer)

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
