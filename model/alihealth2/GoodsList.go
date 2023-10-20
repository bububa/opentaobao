package alihealth2

import (
	"sync"
)

// GoodsList 结构体
type GoodsList struct {
	// 实际价格
	RealPrice string `json:"real_price,omitempty" xml:"real_price,omitempty"`
	// 商品外部编码
	GoodsCode string `json:"goods_code,omitempty" xml:"goods_code,omitempty"`
	// 商品名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 优惠名称
	PromotionName string `json:"promotion_name,omitempty" xml:"promotion_name,omitempty"`
	// 商品数量
	Count int64 `json:"count,omitempty" xml:"count,omitempty"`
	// 0 未优惠 1 已优惠
	PromotionStatus int64 `json:"promotion_status,omitempty" xml:"promotion_status,omitempty"`
	// 商品ID
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 优惠类型
	PromotionType int64 `json:"promotion_type,omitempty" xml:"promotion_type,omitempty"`
	// 是否是优惠商品
	IsPromotion bool `json:"is_promotion,omitempty" xml:"is_promotion,omitempty"`
	// 是否是优惠商品
	Promotion bool `json:"promotion,omitempty" xml:"promotion,omitempty"`
}

var poolGoodsList = sync.Pool{
	New: func() any {
		return new(GoodsList)
	},
}

// GetGoodsList() 从对象池中获取GoodsList
func GetGoodsList() *GoodsList {
	return poolGoodsList.Get().(*GoodsList)
}

// ReleaseGoodsList 释放GoodsList
func ReleaseGoodsList(v *GoodsList) {
	v.RealPrice = ""
	v.GoodsCode = ""
	v.Name = ""
	v.PromotionName = ""
	v.Count = 0
	v.PromotionStatus = 0
	v.Id = 0
	v.PromotionType = 0
	v.IsPromotion = false
	v.Promotion = false
	poolGoodsList.Put(v)
}
