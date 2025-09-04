package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/m16yusuf/belajar-gin/internal/models"
	"github.com/m16yusuf/belajar-gin/internal/utils"
)

type PingHandler struct{}

func NewPingHandler() *PingHandler {
	return &PingHandler{}
}

// GetPing
// @tags 		ping
// @router 	/ping [GET]
// @Param		X-Request-ID 	header 	string 	true 	"Header for send id"
// @Param		Content-Type 	header 	string 	true 	"Header for .."
func (p *PingHandler) GetPing(ctx *gin.Context) {
	requestID := ctx.GetHeader("X-Request-ID")
	contentType := ctx.GetHeader("Content-Type")
	ctx.JSON(http.StatusOK, gin.H{
		"message":     "pong",
		"requestId":   requestID,
		"contentType": contentType,
	})
}

type PingWithParam struct {
	Id     int    `example:"1"`
	Param2 string `example:"yusuf"`
	Q      string `example:"yusufgg"`
}

// GetPingWithParam
// @tags 		ping
// @router 	/ping/:id/:param2 [GET]
// @Param		id			path		integer true 	"path params for id"
// @Param		param2	path		string 	true 	"path params for param2"
// @Param		q				query		string 	false	"query for q"
// @produce json
// @success 200 {object} PingWithParam
func (p *PingHandler) GetPingWithParam(ctx *gin.Context) {
	pingID := ctx.Param("id")
	param2 := ctx.Param("param2")
	q := ctx.Query("q")
	ctx.JSON(http.StatusOK, gin.H{
		"param":  pingID,
		"param2": param2,
		"q":      q,
	})
}

func (p *PingHandler) PostPing(ctx *gin.Context) {
	body := models.Ping{}
	// data-binding, memasukkan body ke dalam variable golang
	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":  err.Error(),
			"succes": false,
		})
		return
	}
	if err := utils.ValidateBody(body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	// log.println(body)
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"body":    body,
		"method":  "POST",
	})
}
