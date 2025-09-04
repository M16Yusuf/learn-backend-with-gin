package middlewares

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func MyLogger(ctx *gin.Context) {
	log.Print("start")
	start := time.Now()
	ctx.Next()
	// next digunakan untuk lanjut ke middleware/handler berikutnya
	duration := time.Since(start)
	log.Printf("durasi request ; %dus\n", duration.Microseconds())
}
