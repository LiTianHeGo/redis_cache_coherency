package rediscachecoherency

//测试数据，商品表
type Product struct {
	Name string `json:"name" gorm:"column:name"`
}

func (Product) TableName() string {
	return "store"
}
