package api

import (
	"net/http"

	db "gita_app/db/sqlc"

	"github.com/gin-gonic/gin"
)

type listChaptersRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=100"`
}

func (server *Server) listChapters(ctx *gin.Context) {
	var req listChaptersRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListChaptersParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	chapters, err := server.store.ListChapters(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, chapters)
}

type addChapterRequest struct {
	Name string `json:"name" binding:"required"`
}

func (server *Server) addChapter(ctx *gin.Context) {
	var req addChapterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	chapter, err := server.store.AddChapter(ctx, req.Name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, chapter)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

type getChapterRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getChapter(ctx *gin.Context) {
	var req getChapterRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	chapter, err := server.store.GetChapter(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, chapter)
}
