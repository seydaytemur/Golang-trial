package main

import (
	"database/sql"
	"deneme/handler"
	"deneme/repository"
	"deneme/service"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:1_Kaankaan@tcp(127.0.0.1:3306)/magazadb")
	if err != nil {
		log.Fatal("Veritabanı bağlantı hatası:", err)
	}
	defer db.Close()

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

	log.Println("Sunucu 8080 portunda başlatıldı.")
	r.Run(":8080")
}
