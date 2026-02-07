// Package main, API uygulamasının giriş noktasıdır.
package main

import (
	"database/sql"
	"deneme/handler"
	"deneme/repository"
	"deneme/service"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	// .env dosyasını yükle
	if err := godotenv.Load(); err != nil {
		log.Fatal("Hata: .env dosyası yüklenemedi. Lütfen .env dosyasını oluşturun:", err)
	}

	// Environment variables'dan veritabanı bilgilerini al
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// Eksik environment variable kontrolü
	if dbUser == "" || dbPassword == "" || dbHost == "" || dbPort == "" || dbName == "" {
		log.Fatal("Hata: Gerekli environment variables eksik. Lütfen .env dosyasını kontrol edin.")
	}

	// Veritabanı bağlantı string'ini oluştur
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Veritabanı bağlantı hatası:", err)
	}
	defer db.Close()

	// Veritabanı bağlantısını test et
	if err := db.Ping(); err != nil {
		log.Fatal("Veritabanı bağlantı testi başarısız:", err)
	}

	urunRepo := repository.NewMySQLUrunRepository(db)
	urunService := service.NewUrunService(urunRepo)
	urunHandler := handler.NewUrunHandler(urunService)

	r := gin.Default()

	r.Static("/static", "./frontend")

	r.GET("/", func(c *gin.Context) {
		htmlFile, err := os.ReadFile("./frontend/index.html")
		if err != nil {
			c.String(http.StatusNotFound, "index.html dosyası bulunamadı.")
			return
		}
		c.Data(http.StatusOK, "text/html", htmlFile)
	})

	r.GET("api/urunler", urunHandler.GetUrunler)
	r.POST("/api/urunler", urunHandler.CreateUrun)
	r.PUT("/api/urunler/:id", urunHandler.UpdateUrun)
	r.DELETE("/api/urunler/:id", urunHandler.DeleteUrun)

	// Port bilgisini environment variable'dan al
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080" // Varsayılan port
	}

	log.Printf("Sunucu %s portunda başlatıldı.", port)
	r.Run(":" + port)
}
