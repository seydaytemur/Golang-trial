package service

import (
	"deneme/models"
	"deneme/repository"
)

// UrunService, ürün işlemlerinin iş mantığını sağlar.
type UrunService struct {
	repo repository.UrunRepository
}

// NewUrunService, yeni bir UrunService oluşturur.
func NewUrunService(r repository.UrunRepository) *UrunService {
	return &UrunService{repo: r}
}

// GetAllUrunler, tüm ürünleri getirir.
func (s *UrunService) GetAllUrunler() ([]models.Urun, error) {
	return s.repo.GetAll()
}

// CreateUrun, yeni bir ürün ekler.
func (s *UrunService) CreateUrun(urun models.Urun) (int64, error) {

	return s.repo.Create(urun)
}

// UpdateUrun, ürünü günceller.
func (s *UrunService) UpdateUrun(urun models.Urun) (int64, error) {

	return s.repo.Update(urun)
}

// DeleteUrun, ürünü siler.
func (s *UrunService) DeleteUrun(id int) (int64, error) {

	return s.repo.Delete(id)
}
