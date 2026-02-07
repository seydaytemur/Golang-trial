package handler

import (
	"bytes"
	"deneme/models"
	"deneme/service"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// Mock Repository for Handler Tests
type MockRepo struct {
	MockGetAll func() ([]models.Urun, error)
	MockCreate func(urun models.Urun) (int64, error)
	MockUpdate func(urun models.Urun) (int64, error)
	MockDelete func(id int) (int64, error)
}

func (m *MockRepo) GetAll() ([]models.Urun, error) {
	if m.MockGetAll != nil {
		return m.MockGetAll()
	}
	return nil, nil
}

func (m *MockRepo) Create(urun models.Urun) (int64, error) {
	if m.MockCreate != nil {
		return m.MockCreate(urun)
	}
	return 0, nil
}

func (m *MockRepo) Update(urun models.Urun) (int64, error) {
	if m.MockUpdate != nil {
		return m.MockUpdate(urun)
	}
	return 0, nil
}

func (m *MockRepo) Delete(id int) (int64, error) {
	if m.MockDelete != nil {
		return m.MockDelete(id)
	}
	return 0, nil
}

func TestGetUrunler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockRepo := &MockRepo{
		MockGetAll: func() ([]models.Urun, error) {
			return []models.Urun{
				{ID: 1, UrunAdi: "Test", StokMiktari: 10, Fiyat: 50.0},
			}, nil
		},
	}

	svc := service.NewUrunService(mockRepo)
	handler := NewUrunHandler(svc)

	r := gin.Default()
	r.GET("/api/urunler", handler.GetUrunler)

	req, _ := http.NewRequest("GET", "/api/urunler", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response []models.Urun
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(response))
	assert.Equal(t, "Test", response[0].UrunAdi)
}

func TestCreateUrun(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockRepo := &MockRepo{
		MockCreate: func(u models.Urun) (int64, error) {
			return 1, nil
		},
	}

	svc := service.NewUrunService(mockRepo)
	handler := NewUrunHandler(svc)

	r := gin.Default()
	r.POST("/api/urunler", handler.CreateUrun)

	urun := models.Urun{UrunAdi: "Yeni", StokMiktari: 5, Fiyat: 10.0}
	jsonValue, _ := json.Marshal(urun)

	req, _ := http.NewRequest("POST", "/api/urunler", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestCreateUrun_Error(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockRepo := &MockRepo{
		MockCreate: func(u models.Urun) (int64, error) {
			return 0, errors.New("db error")
		},
	}

	svc := service.NewUrunService(mockRepo)
	handler := NewUrunHandler(svc)

	r := gin.Default()
	r.POST("/api/urunler", handler.CreateUrun)

	urun := models.Urun{UrunAdi: "Hatalı", StokMiktari: 5, Fiyat: 10.0}
	jsonValue, _ := json.Marshal(urun)

	req, _ := http.NewRequest("POST", "/api/urunler", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}
