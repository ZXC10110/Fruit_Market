package Model

import (
	"github.com/zxc10110/Fruit_Market/Config"

	_ "github.com/go-sql-driver/mysql"
)

//GetAllHospitals
func GetAllFruit() (fruit []Fruit, err error) {
	if err = Config.DB.Find(&fruit).Error; err != nil {
		return
	}
	return
}
