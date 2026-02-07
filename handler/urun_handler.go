package handler

import (
	"deneme/models"
	"deneme/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
)

// UrunHandler, ürünlerle ilgili HTTP isteklerini yönetir.
type UrunHandler struct {
	service *service.UrunService
}

// NewUrunHandler, yeni bir UrunHandler örneği oluşturur.
func NewUrunHandler(s *service.UrunService) *UrunHandler {
	return &UrunHandler{service: s}
}

// GetUrunler, tüm ürünleri listeler.
func (h *UrunHandler) GetUrunler(c *gin.Context) {
	urunler, err := h.service.GetAllUrunler()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ürünler getirilemedi"})
		return
	}
	c.JSON(http.StatusOK, urunler)
}

// CreateUrun, yeni bir ürün oluşturur.
func (h *UrunHandler) CreateUrun(c *gin.Context) {
	var urun models.Urun
	if err := c.BindJSON(&urun); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Geçersiz veri"})
		return
	}
	id, err := h.service.CreateUrun(urun)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ürün oluşturulamadı"})
		return
	}
	urun.ID = int(id)
	c.JSON(http.StatusCreated, urun)
}

// UpdateUrun, mevcut bir ürünü günceller.
func (h *UrunHandler) UpdateUrun(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var urun models.Urun
	if err := c.BindJSON(&urun); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Geçersiz veri"})
		return
	}
	urun.ID = id

	_, err := h.service.UpdateUrun(urun)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ürün güncellenemedi"})
		return
	}
	c.JSON(http.StatusOK, urun)
}

// DeleteUrun, bir ürünü siler.
func (h *UrunHandler) DeleteUrun(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	_, err := h.service.DeleteUrun(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ürün silinemedi"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Ürün başarıyla silindi"})
}
