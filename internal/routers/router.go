package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/m16yusuf/belajar-gin/internal/middlewares"
	"github.com/m16yusuf/belajar-gin/internal/models"
)

func InitRouter(db *pgxpool.Pool) *gin.Engine {
	// inisialisasi engine gin
	router := gin.Default()
	router.Use(middlewares.MyLogger)
	router.Use(middlewares.CORSMiddleware)

	// setup routing
	InitPingRouter(router)
	InitRentalRouter(router, db)

	// jika route tidak ditemukan kirim response
	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, models.Response{
			Message: "salah...",
			Status:  "Tidak ditemukan",
		})
	})

	return router
}
