package service 

import (
	"deneme/models"
	"deneme/repository"
)

type UrunService struct {
	repo repository.UrunRepository
}

func NewUrunService(r repository.UrunRepository) *UrunService {
	return &UrunService{repo: r}
}

func (s *UrunService) GetAllUrunler() ([]models.Urun, error) {
	return s.repo.GetAll()
}

func (s *UrunService) CreateUrun(urun models.Urun) (int64, error) {
	
	return s.repo.Create(urun)
}
func (s *UrunService) UpdateUrun(urun models.Urun) (int64, error) {
	
	return s.repo.Update(urun)
}

func (s *UrunService) DeleteUrun(id int) (int64, error) {
	
	return s.repo.Delete(id)
}
