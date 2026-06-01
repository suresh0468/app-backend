package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

type getSlokaRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getSloka(ctx *gin.Context) {
	var req getSlokaRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	sloka, err := server.store.GetSloka(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, sloka)
}

type listSlokasByChapterRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) listSlokasByChapter(ctx *gin.Context) {
	var req listSlokasByChapterRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	slokas, err := server.store.ListSlokasByChapter(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, slokas)
}
