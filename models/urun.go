package models

type Urun struct {
	ID          int     `json:"id"`
	UrunAdi     string  `json:"urun_adi"`
	StokMiktari int     `json:"stok_miktari"`
	Fiyat       float64 `json:"fiyat"`
}
