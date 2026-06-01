package api

import (
	db "gita_app/db/sqlc"

	"github.com/gin-gonic/gin"
)

type Server struct {
	store  *db.Queries
	router *gin.Engine
}

func NewServer(store *db.Queries) *Server {
	server := &Server{
		store: store,
	}

	router := gin.Default()

	router.POST("/chapters", server.addChapter)
	router.GET("/chapters", server.listChapters)
	router.GET("/chapters/:id/slokas", server.listSlokasByChapter)
	router.GET("/slokas/:id", server.getSloka)

	server.router = router

	return server
}

func (s *Server) Start(address string) error {
	return s.router.Run(address)
}
