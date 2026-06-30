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
	Title       string `json:"title" binding:"required"`
	Subtitle    string `json:"subtitle"`
	Description string `json:"description"`
}

func (server *Server) addChapter(ctx *gin.Context) {
	var req addChapterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.AddChapterParams{
		Title:       req.Title,
		Subtitle:    req.Subtitle,
		Description: req.Description,
	}

	chapter, err := server.store.AddChapter(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, chapter)
}

func (server *Server) addChapters(ctx *gin.Context) {
	var req []addChapterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	var chapters []db.Chapter
	for _, reqChapter := range req {
		arg := db.AddChapterParams{
			Title:       reqChapter.Title,
			Subtitle:    reqChapter.Subtitle,
			Description: reqChapter.Description,
		}

		chapter, err := server.store.AddChapter(ctx, arg)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		chapters = append(chapters, chapter)
	}

	ctx.JSON(http.StatusOK, chapters)
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

type updateChapterUriRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

type updateChapterBodyRequest struct {
	Title       string `json:"title"`
	Subtitle    string `json:"subtitle"`
	Description string `json:"description"`
}

func (server *Server) updateChapter(ctx *gin.Context) {
	var uri updateChapterUriRequest
	if err := ctx.ShouldBindUri(&uri); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	var body updateChapterBodyRequest
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// Fetch the current record so unchanged fields stay intact via COALESCE.
	existing, err := server.store.GetChapter(ctx, uri.ID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, errorResponse(err))
		return
	}

	arg := db.UpdateChapterParams{
		ID:          uri.ID,
		Title:       existing.Title,
		Subtitle:    existing.Subtitle,
		Description: existing.Description,
	}
	if body.Title != "" {
		arg.Title = body.Title
	}
	if body.Subtitle != "" {
		arg.Subtitle = body.Subtitle
	}
	if body.Description != "" {
		arg.Description = body.Description
	}

	chapter, err := server.store.UpdateChapter(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, chapter)
}
