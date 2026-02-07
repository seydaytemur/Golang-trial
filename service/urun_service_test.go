package service

import (
	"deneme/models"
	"errors"
	"testing"
)

func TestGetAllUrunler(t *testing.T) {
	expectedUrunler := []models.Urun{
		{ID: 1, UrunAdi: "Test Urun 1", StokMiktari: 10, Fiyat: 100.0},
		{ID: 2, UrunAdi: "Test Urun 2", StokMiktari: 20, Fiyat: 200.0},
	}

	mockRepo := &MockUrunRepository{
		MockGetAll: func() ([]models.Urun, error) {
			return expectedUrunler, nil
		},
	}

	service := NewUrunService(mockRepo)

	urunler, err := service.GetAllUrunler()

	if err != nil {
		t.Fatalf("Beklenmedik hata: %v", err)
	}

	if len(urunler) != len(expectedUrunler) {
		t.Errorf("Beklenen ürün sayısı %d, fakat %d alındı", len(expectedUrunler), len(urunler))
	}

	if urunler[0].UrunAdi != expectedUrunler[0].UrunAdi {
		t.Errorf("Beklenen ürün adı %s, fakat %s alındı", expectedUrunler[0].UrunAdi, urunler[0].UrunAdi)
	}
}

func TestGetAllUrunler_Error(t *testing.T) {
	mockRepo := &MockUrunRepository{
		MockGetAll: func() ([]models.Urun, error) {
			return nil, errors.New("veritabanı hatası")
		},
	}

	service := NewUrunService(mockRepo)
	_, err := service.GetAllUrunler()

	if err == nil {
		t.Error("Hata bekleniyordu fakat nil döndü")
	}
}

func TestCreateUrun(t *testing.T) {
	urun := models.Urun{UrunAdi: "Yeni Urun", StokMiktari: 5, Fiyat: 50.0}

	mockRepo := &MockUrunRepository{
		MockCreate: func(u models.Urun) (int64, error) {
			return 1, nil
		},
	}

	service := NewUrunService(mockRepo)
	id, err := service.CreateUrun(urun)

	if err != nil {
		t.Fatalf("CreateUrun sırasında hata: %v", err)
	}

	if id != 1 {
		t.Errorf("Beklenen ID 1, alınan %d", id)
	}
}

func TestDeleteUrun(t *testing.T) {
	mockRepo := &MockUrunRepository{
		MockDelete: func(id int) (int64, error) {
			if id == 1 {
				return 1, nil
			}
			return 0, errors.New("bulunamadı")
		},
	}

	service := NewUrunService(mockRepo)

	// Başarılı silme
	rowsAffected, err := service.DeleteUrun(1)
	if err != nil {
		t.Errorf("Silme sırasında beklenmedik hata: %v", err)
	}
	if rowsAffected != 1 {
		t.Errorf("Etkilenen satır 1 olmalıydı, alınan %d", rowsAffected)
	}

	// Başarısız silme
	_, err = service.DeleteUrun(99)
	if err == nil {
		t.Error("Var olmayan ID için hata bekleniyordu")
	}
}
