package nlife

import (
	"sync"
)

// Goods 结构体
type Goods struct {
	// 导购员编号
	Guider string `json:"guider,omitempty" xml:"guider,omitempty"`
	// 币种
	Currency string `json:"currency,omitempty" xml:"currency,omitempty"`
	// 零售+商品ID
	Id string `json:"id,omitempty" xml:"id,omitempty"`
	// 商品标题
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 商家自用编码集合，对应该item_sku. 逗号分隔字符串
	CustomCodes string `json:"custom_codes,omitempty" xml:"custom_codes,omitempty"`
	// 商品数量
	Quantity int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// 商品价格，人民币 分
	Price int64 `json:"price,omitempty" xml:"price,omitempty"`
	// 商品itemId
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 商品skuId
	SkuId int64 `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
}

var poolGoods = sync.Pool{
	New: func() any {
		return new(Goods)
	},
}

// GetGoods() 从对象池中获取Goods
func GetGoods() *Goods {
	return poolGoods.Get().(*Goods)
}

// ReleaseGoods 释放Goods
func ReleaseGoods(v *Goods) {
	v.Guider = ""
	v.Currency = ""
	v.Id = ""
	v.Title = ""
	v.CustomCodes = ""
	v.Quantity = 0
	v.Price = 0
	v.ItemId = 0
	v.SkuId = 0
	poolGoods.Put(v)
}
