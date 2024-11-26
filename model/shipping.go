package model

type Shipping struct {
	Id   int
	Name string
	// Table string `gorm:"table:Shipping"`
}

func (Shipping) TableName() string {
	return "Shipping" // Nama tabel yang diinginkan adalah 'shipping'
}

type Ongkir struct {
	IdShipping         int
	Qty                int
	LatLongOrigin      string
	LatLongDestination string
	Distance           float64
	Price              int
}
