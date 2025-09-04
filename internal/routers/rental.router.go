package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/m16yusuf/belajar-gin/internal/handlers"
	"github.com/m16yusuf/belajar-gin/internal/repositories"
)

func InitRentalRouter(router *gin.Engine, db *pgxpool.Pool) {
	rentalRouter := router.Group("/rentals")
	rentalRepository := repositories.NewRentalRepository(db)
	rh := handlers.NewRentalHandler(rentalRepository)

	rentalRouter.GET("", rh.GetRental)
	rentalRouter.POST("", rh.PostRental)
	rentalRouter.PATCH("/:id", rh.PatchRental)
}
