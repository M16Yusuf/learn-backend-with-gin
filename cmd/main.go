package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/m16yusuf/belajar-gin/internal/configs"
	"github.com/m16yusuf/belajar-gin/internal/routers"
)

func main() {
	// manual load ENV
	if err := godotenv.Load(); err != nil {
		log.Println("Failed to load env \nCause: ", err.Error())
		return
	}

	// inisialisasi database
	db, err := configs.InitDB()
	if err != nil {
		log.Println("Failed to connect to  database\nCause: ", err.Error())
		return
	}
	defer db.Close()
	// testing koneksi db
	if err := configs.TestDB(db); err != nil {
		log.Println("Ping to DB failsed\nCause: ", err.Error())
	}
	log.Println("DB connected")

	// inisialisasi engine gin
	router := routers.InitRouter(db)
	// Jalankan engine gin
	// default : localhost/0.0.0.0/127.0.0.1 : 8080
	router.Run()
}
