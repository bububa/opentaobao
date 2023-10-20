package eleenterpriseordernew

import (
	"sync"
)

// FoodsInfo 结构体
type FoodsInfo struct {
	// 餐品名称
	FoodName string `json:"food_name,omitempty" xml:"food_name,omitempty"`
	// 餐品价格
	Price string `json:"price,omitempty" xml:"price,omitempty"`
	// 餐品id
	FoodId int64 `json:"food_id,omitempty" xml:"food_id,omitempty"`
	// 餐品数量
	Count int64 `json:"count,omitempty" xml:"count,omitempty"`
	// 规格Id
	SkuId int64 `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
}

var poolFoodsInfo = sync.Pool{
	New: func() any {
		return new(FoodsInfo)
	},
}

// GetFoodsInfo() 从对象池中获取FoodsInfo
func GetFoodsInfo() *FoodsInfo {
	return poolFoodsInfo.Get().(*FoodsInfo)
}

// ReleaseFoodsInfo 释放FoodsInfo
func ReleaseFoodsInfo(v *FoodsInfo) {
	v.FoodName = ""
	v.Price = ""
	v.FoodId = 0
	v.Count = 0
	v.SkuId = 0
	poolFoodsInfo.Put(v)
}
