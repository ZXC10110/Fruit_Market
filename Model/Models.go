package Model

func (b *Fruit) TableName() string {
	return "fruit"
}

type Fruit struct {
	Fid    int     `json:"f_id"`
	FName  string  `json:"f_name"`
	FPrice float64 `json:"f_price"`
	FQty   int     `json:"f_qty"`
}
