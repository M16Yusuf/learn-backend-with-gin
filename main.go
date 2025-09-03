package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"slices"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Rental struct {
	Id         int        `db:"id" json:"id"`
	Image      string     `db:"image" json:"image_path"`
	Name       string     `db:"rentals_name" json:"rental_name"`
	User_id    int        `db:"user_id" json:"user_id"`
	Created_at time.Time  `db:"created_at" json:"created_at"`
	Updated_at *time.Time `db:"updated_at" json:"updated_at"`
}

func main() {
	// manual load
	if err := godotenv.Load(); err != nil {
		log.Println("Failed to load env \nCause: ", err.Error())
		return
	}

	// pakai env
	// log.Println(os.Getenv("DB_YUSUF"))

	// data database berada di .env
	dbUser := os.Getenv("DB_USER")
	dbUserPass := os.Getenv("DB_USER_PASS")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	// inisilisasi DB
	connetString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbUserPass, dbHost, dbPort, dbName)
	db, err := pgxpool.New(context.Background(), connetString)
	if err != nil {
		log.Println("Failed to connect to  database\nCause: ", err.Error())
		return
	}
	defer db.Close()

	// testing koneksi db
	if err := db.Ping(context.Background()); err != nil {
		log.Println("Ping to DB failsed\nCause: ", err.Error())
	}
	log.Println("DB connected")

	// inisialisasi engine gin
	router := gin.Default()
	router.Use(MyLogger)
	router.Use(CORSMiddleware)

	// memakai database pada https method
	router.GET("/rentals", func(ctx *gin.Context) {
		// membuat pagenation menggunakan query LIMIT dan OFFSET
		page, err := strconv.Atoi(ctx.Query("page"))
		if err != nil {
			page = 1
		}
		limit := 4
		offset := (page - 1) * limit
		// query ke database
		sql := "SELECT * FROM rentals LIMIT $2 OFFSET $1"
		values := []any{offset, limit}
		rows, err := db.Query(ctx.Request.Context(), sql, values...)
		if err != nil {
			log.Println("internal server error : ", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"data":    []any{},
			})
			return
		}
		defer rows.Close()

		// mengolah data / membaca rows/record
		var rentals []Rental
		for rows.Next() {
			var rental Rental
			if err := rows.Scan(&rental.Id, &rental.Image, &rental.Name, &rental.User_id, &rental.Created_at, &rental.Updated_at); err != nil {
				log.Println("Scan Error, ", err.Error())
				return
			}
			rentals = append(rentals, rental)
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

	})

	// melakukan Insert ke database dengan data dari body https
	router.POST("/rentals", func(ctx *gin.Context) {
		var body Rental
		if err := ctx.ShouldBind(&body); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error":  err.Error(),
				"succes": false,
			})
			return
		}
		sql := "INSERT INTO rentals (image, rentals_name, user_id) VALUES ($1, $2, $3) RETURNING id, image, rentals_name, user_id, created_at"
		values := []any{body.Image, body.Name, body.User_id}
		var newRental Rental
		if err := db.QueryRow(ctx.Request.Context(), sql, values...).Scan(&newRental.Id, &newRental.Image, &newRental.Name, &newRental.User_id, &newRental.Created_at); err != nil {
			log.Println("Scan Error, ", err.Error())
			return
		}

		ctx.JSON(http.StatusCreated, gin.H{
			"success": true,
			"data":    newRental,
		})
	})

	// melakukuan update denge method patch, hanya berubah sebagian
	// data yang mungkin bisa dirubah : image, rentals_name, user_id
	router.PATCH("/rentals", func(ctx *gin.Context) {
		// binding data dengen model Rental
		var body Rental
		if err := ctx.ShouldBind(&body); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"errors":  err.Error(),
				"success": false,
			})
		}

		// query dan pengecekan apa saja yang akan dirubah
		sql := "UPDATE rentals SET updated_at=CURRENT_TIMESTAMP "
		if body.Image != "" {
			sql += ", image='" + body.Image + "'"
		}
		if body.Name != "" {
			sql += ", rentals_name='" + body.Name + "'"
		}
		if body.User_id != 0 {
			sql += ", user_id=" + strconv.Itoa(body.User_id) + ""
		}
		sql += " WHERE id=$1 RETURNING id, image, rentals_name, user_id, created_at, updated_at "
		values := []any{body.Id}

		// masukan data hasil returning query ke dalam model rental yang baru
		var newRental Rental
		if err := db.QueryRow(ctx.Request.Context(), sql, values...).Scan(&newRental.Id, &newRental.Image, &newRental.Name, &newRental.User_id, &newRental.Created_at, &newRental.Updated_at); err != nil {
			log.Println("scan Error. ", err.Error())
			return
		}

		// return model data baru rental sebagai respons
		ctx.JSON(http.StatusOK, gin.H{
			"success": true,
			"data":    newRental,
		})

	})

	// setup rrouting
	router.GET("/ping", func(ctx *gin.Context) {
		requestID := ctx.GetHeader("X-Request-ID")
		contentType := ctx.GetHeader("Content-Type")
		ctx.JSON(http.StatusOK, gin.H{
			"message":     "pong",
			"requestId":   requestID,
			"contentType": contentType,
		})
	})

	router.GET("/ping/:id/:param2", func(ctx *gin.Context) {
		pingID := ctx.Param("id")
		param2 := ctx.Param("param2")
		q := ctx.Query("q")
		ctx.JSON(http.StatusOK, gin.H{
			"param":  pingID,
			"param2": param2,
			"q":      q,
		})
	})

	router.POST("/ping", func(ctx *gin.Context) {
		body := Body{}
		// data-binding, memasukkan body ke dalam variable golang
		if err := ctx.ShouldBind(&body); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error":  err.Error(),
				"succes": false,
			})
			return
		}
		if err := ValidateBody(body); err != nil {
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
	})

	router.PATCH("/ping", func(ctx *gin.Context) {
		body := Body{}
		// data-binding, memasukkan body ke dalam variable golang
		if err := ctx.ShouldBind(&body); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error":  err.Error(),
				"succes": false,
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"success": true,
			"body":    body,
			"method":  "PATCH",
		})
	})

	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, Response{
			Message: "salah...",
			Status:  "Tidak ditemukan",
		})
	})
	router.Run()
}

func MyLogger(ctx *gin.Context) {
	log.Print("start")
	start := time.Now()
	ctx.Next()
	// next digunakan untuk lanjut ke middleware/handler berikutnya
	duration := time.Since(start)
	log.Printf("durasi request ; %dus\n", duration.Microseconds())
}

func CORSMiddleware(ctx *gin.Context) {
	// memasangkan header-header CORS
	// setup whitelist origin
	whitelist := []string{"http://127.0.0.1:5500"}
	origin := ctx.GetHeader("Origin")
	if slices.Contains(whitelist, origin) {
		ctx.Header("Access-Control-Allow-Origin", origin)
	}
	// header untuk preflight cors
	ctx.Header("Access-Control-Allow-Methods", "GET")
	ctx.Header("Access-Control-Allow-Headers", "Authorization, Content-Type")
	// tangani apabila bertemu preflight
	if ctx.Request.Method == http.MethodOptions {
		ctx.AbortWithStatus(http.StatusNoContent)
		return
	}
	ctx.Next()
}

type Response struct {
	Message string
	Status  string
}

type Body struct {
	Id      int    `json:"id"`
	Message string `json:"msg"`
	Gender  string `json:"gender"`
}

func ValidateBody(body Body) error {
	if body.Id <= 0 {
		return errors.New("id tidak boleh dibawah 0")
	}
	if len(body.Message) < 8 {
		return errors.New("panjang pesan harus diatas 8 karakter")
	}
	re, err := regexp.Compile("^[lLpPmMfF]$")
	if err != nil {
		return err
	}
	if isMatched := re.Match([]byte(body.Gender)); !isMatched {
		return errors.New("gender harus berisikan huruf l, L, m, M, f, F, p, P")
	}
	return nil
}
