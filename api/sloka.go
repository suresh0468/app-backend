package api

import (
	"database/sql"
	"net/http"

	db "gita_app/db/sqlc"

	"github.com/gin-gonic/gin"
)

type slokaResponse struct {
	ID              int64  `json:"id"`
	ChapterID       int64  `json:"chapter_id"`
	Sloka           string `json:"sloka"`
	Transliteration string `json:"transliteration"`
	Purport         string `json:"purport,omitempty"`
	Explanation     string `json:"explanation,omitempty"`
}

func newSlokaResponse(sloka db.Sloka) slokaResponse {
	return slokaResponse{
		ID:              sloka.ID,
		ChapterID:       sloka.ChapterID,
		Sloka:           sloka.Sloka,
		Transliteration: sloka.Transliteration,
		Purport:         sloka.Purport.String,
		Explanation:     sloka.Explanation.String,
	}
}

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

	ctx.JSON(http.StatusOK, newSlokaResponse(sloka))
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

	rsp := make([]slokaResponse, len(slokas))
	for i, sloka := range slokas {
		rsp[i] = newSlokaResponse(sloka)
	}

	ctx.JSON(http.StatusOK, rsp)
}

type addSlokaRequest struct {
	ChapterID       int64  `json:"chapter_id" binding:"required,min=1"`
	Sloka           string `json:"sloka" binding:"required"`
	Transliteration string `json:"transliteration" binding:"required"`
	Purport         string `json:"purport"`
	Explanation     string `json:"explanation"`
}

func (server *Server) addSloka(ctx *gin.Context) {
	var req addSlokaRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.AddSlokaParams{
		ChapterID:       req.ChapterID,
		Sloka:           req.Sloka,
		Transliteration: req.Transliteration,
		Purport:         sql.NullString{String: req.Purport, Valid: req.Purport != ""},
		Explanation:     sql.NullString{String: req.Explanation, Valid: req.Explanation != ""},
	}

	sloka, err := server.store.AddSloka(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, newSlokaResponse(sloka))
}

type updateSlokaUriRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

type updateSlokaBodyRequest struct {
	Sloka           string `json:"sloka"`
	Transliteration string `json:"transliteration"`
	Purport         string `json:"purport"`
	Explanation     string `json:"explanation"`
}

func (server *Server) updateSloka(ctx *gin.Context) {
	var uri updateSlokaUriRequest
	if err := ctx.ShouldBindUri(&uri); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	var body updateSlokaBodyRequest
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// Fetch the current record so unchanged fields stay intact.
	existing, err := server.store.GetSloka(ctx, uri.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	arg := db.UpdateSlokaParams{
		ID:              uri.ID,
		Sloka:           existing.Sloka,
		Transliteration: existing.Transliteration,
		Purport:         existing.Purport,
		Explanation:     existing.Explanation,
	}
	if body.Sloka != "" {
		arg.Sloka = body.Sloka
	}
	if body.Transliteration != "" {
		arg.Transliteration = body.Transliteration
	}
	if body.Purport != "" {
		arg.Purport = sql.NullString{String: body.Purport, Valid: true}
	}
	if body.Explanation != "" {
		arg.Explanation = sql.NullString{String: body.Explanation, Valid: true}
	}

	sloka, err := server.store.UpdateSloka(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, newSlokaResponse(sloka))
}
