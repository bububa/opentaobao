package icbulogistics

import (
	"sync"
)

// CargoList 结构体
type CargoList struct {
	// 商品特性列表对象
	ProductType []ProductType `json:"product_type,omitempty" xml:"product_type>product_type,omitempty"`
	// 单位
	Unit string `json:"unit,omitempty" xml:"unit,omitempty"`
	// 海关编码
	Hscode string `json:"hscode,omitempty" xml:"hscode,omitempty"`
	// 1
	DeclarationValue string `json:"declaration_value,omitempty" xml:"declaration_value,omitempty"`
	// 货物单价
	Price string `json:"price,omitempty" xml:"price,omitempty"`
	// 货物中文名
	NameCn string `json:"name_cn,omitempty" xml:"name_cn,omitempty"`
	// 1
	Currency string `json:"currency,omitempty" xml:"currency,omitempty"`
	// 货物英文名
	NameEn string `json:"name_en,omitempty" xml:"name_en,omitempty"`
	// 材质
	Material string `json:"material,omitempty" xml:"material,omitempty"`
	// 用途
	Purpose string `json:"purpose,omitempty" xml:"purpose,omitempty"`
	// 货物数量
	Quantity int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
}

var poolCargoList = sync.Pool{
	New: func() any {
		return new(CargoList)
	},
}

// GetCargoList() 从对象池中获取CargoList
func GetCargoList() *CargoList {
	return poolCargoList.Get().(*CargoList)
}

// ReleaseCargoList 释放CargoList
func ReleaseCargoList(v *CargoList) {
	v.ProductType = v.ProductType[:0]
	v.Unit = ""
	v.Hscode = ""
	v.DeclarationValue = ""
	v.Price = ""
	v.NameCn = ""
	v.Currency = ""
	v.NameEn = ""
	v.Material = ""
	v.Purpose = ""
	v.Quantity = 0
	poolCargoList.Put(v)
}
