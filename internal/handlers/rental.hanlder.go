package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/m16yusuf/belajar-gin/internal/models"
	"github.com/m16yusuf/belajar-gin/internal/repositories"
)

type RentalHandler struct {
	renRepo *repositories.RentalRepository
}

func NewRentalHandler(renRepo *repositories.RentalRepository) *RentalHandler {
	return &RentalHandler{
		renRepo: renRepo,
	}
}

// get method http untuk mendapatkan data rentals dari database
func (r *RentalHandler) GetRental(ctx *gin.Context) {
	// membuat pagenation menggunakan query LIMIT dan OFFSET
	page, err := strconv.Atoi(ctx.Query("page"))
	if err != nil {
		page = 1
	}
	limit := 4
	offset := (page - 1) * limit

	rentals, err := r.renRepo.GetRental(ctx.Request.Context(), offset, limit)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"data":    rentals,
		})
		return
	}

	// error handling jika data kosong
	if len(rentals) == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"data":    []any{},
			"page":    page,
		})
		return
	}

	// Return data sebagai respons dengan data
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    rentals,
		"page":    page,
	})
}

// melakukan Insert ke database dengan data dari body https
func (r *RentalHandler) PostRental(ctx *gin.Context) {
	var body models.Rental
	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":  err.Error(),
			"succes": false,
		})
		return
	}

	newRental, err := r.renRepo.NewRental(ctx.Request.Context(), body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    newRental,
	})
}

// melakukuan update denge method patch, hanya berubah sebagian
// data yang mungkin bisa dirubah : image, rentals_name, user_id
func (r *RentalHandler) PatchRental(ctx *gin.Context) {
	// binding data dengen model Rental
	var body models.Rental
	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"errors":  err.Error(),
			"success": false,
		})
		return
	}

	paramId := ctx.Param("id")
	newRental, err := r.renRepo.UpdateRental(ctx.Request.Context(), body, paramId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
		})
		return
	}

	// return model data baru rental sebagai respons
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    newRental,
	})
}
