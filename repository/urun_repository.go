package repository



import (
	"database/sql"
	"deneme/models"
	"log"
)

type UrunRepository interface {
	GetAll() ([]models.Urun, error)
	Create(urun models.Urun) (int64, error)
	Update(urun models.Urun) (int64, error)
	Delete(id int) (int64, error)
}

type mysqlUrunRepository struct {
	db *sql.DB
}

func NewMySQLUrunRepository(db *sql.DB) UrunRepository {
	return &mysqlUrunRepository{db: db}
}

func (r *mysqlUrunRepository) GetAll() ([]models.Urun, error) {
	rows, err := r.db.Query("SELECT id, urun_adi, stok_miktari, fiyat FROM urunler")
	if err != nil {
		log.Fatal("GetAll sorgusunda hata:", err)
		return nil, err
	}
	defer rows.Close()

	var urunler []models.Urun
	for rows.Next() {
		var urun models.Urun
		
		if err := rows.Scan(&urun.ID, &urun.UrunAdi, &urun.StokMiktari, &urun.Fiyat); err != nil {
			log.Fatal("GetAll satır tarama (rows.Scan) hatası:", err) // Hata ayıklama satırı
			return nil, err
		}
		urunler = append(urunler, urun)
	}

	
	if err := rows.Err(); err != nil {
		log.Fatal("GetAll satır hatası:", err) 
		return nil, err
	}

	return urunler, nil
}

func (r *mysqlUrunRepository) Create(urun models.Urun) (int64, error) {
	
	res, err := r.db.Exec("INSERT INTO urunler (urun_adi, stok_miktari, fiyat) VALUES (?, ?, ?)", urun.UrunAdi, urun.StokMiktari, urun.Fiyat)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func (r *mysqlUrunRepository) Update(urun models.Urun) (int64, error) {
	
	res, err := r.db.Exec("UPDATE urunler SET urun_adi=?, stok_miktari=?, fiyat=? WHERE id=? ", urun.UrunAdi, urun.StokMiktari, urun.Fiyat, urun.ID)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func (r *mysqlUrunRepository) Delete(id int) (int64, error) {
	res, err := r.db.Exec("DELETE FROM urunler WHERE id = ?", id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}
