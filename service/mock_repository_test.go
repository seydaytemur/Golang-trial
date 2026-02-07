package service

import (
	"deneme/models"
)

type MockUrunRepository struct {
	MockGetAll func() ([]models.Urun, error)
	MockCreate func(urun models.Urun) (int64, error)
	MockUpdate func(urun models.Urun) (int64, error)
	MockDelete func(id int) (int64, error)
}

func (m *MockUrunRepository) GetAll() ([]models.Urun, error) {
	if m.MockGetAll != nil {
		return m.MockGetAll()
	}
	return nil, nil
}

func (m *MockUrunRepository) Create(urun models.Urun) (int64, error) {
	if m.MockCreate != nil {
		return m.MockCreate(urun)
	}
	return 0, nil
}

func (m *MockUrunRepository) Update(urun models.Urun) (int64, error) {
	if m.MockUpdate != nil {
		return m.MockUpdate(urun)
	}
	return 0, nil
}

func (m *MockUrunRepository) Delete(id int) (int64, error) {
	if m.MockDelete != nil {
		return m.MockDelete(id)
	}
	return 0, nil
}
