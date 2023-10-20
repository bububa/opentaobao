package eleenterpriserestaurant

import (
	"sync"
)

// Specfood 结构体
type Specfood struct {
	// 规格信息（示例为多规格例子，否则为 []）
	Specs []Spec `json:"specs,omitempty" xml:"specs>spec,omitempty"`
	// 名字
	FoodName string `json:"food_name,omitempty" xml:"food_name,omitempty"`
	// 原价, 如果没有绑定活动, 该字段为空
	OriginalPrice string `json:"original_price,omitempty" xml:"original_price,omitempty"`
	// 当前价格, 如果美食活动, 此价格是优惠后的价格
	Price string `json:"price,omitempty" xml:"price,omitempty"`
	// 食物id
	FoodId string `json:"food_id,omitempty" xml:"food_id,omitempty"`
	// 商品规格id
	SkuId string `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	// 库存
	Stock int64 `json:"stock,omitempty" xml:"stock,omitempty"`
	// 是否必点, 必点分类下只要点一个 food 就好了, 如果必点商品售完, 不需要必点
	IsEssential bool `json:"is_essential,omitempty" xml:"is_essential,omitempty"`
	// 是否售完
	SoldOut bool `json:"sold_out,omitempty" xml:"sold_out,omitempty"`
}

var poolSpecfood = sync.Pool{
	New: func() any {
		return new(Specfood)
	},
}

// GetSpecfood() 从对象池中获取Specfood
func GetSpecfood() *Specfood {
	return poolSpecfood.Get().(*Specfood)
}

// ReleaseSpecfood 释放Specfood
func ReleaseSpecfood(v *Specfood) {
	v.Specs = v.Specs[:0]
	v.FoodName = ""
	v.OriginalPrice = ""
	v.Price = ""
	v.FoodId = ""
	v.SkuId = ""
	v.Stock = 0
	v.IsEssential = false
	v.SoldOut = false
	poolSpecfood.Put(v)
}
