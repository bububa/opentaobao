package trade

import (
	"sync"
)

// GoodsDetail 结构体
type GoodsDetail struct {
	// 商品标识
	GoodsId string `json:"goods_id,omitempty" xml:"goods_id,omitempty"`
	// 商品单价，人民币：分
	Price string `json:"price,omitempty" xml:"price,omitempty"`
	// 商品数量
	Quantity string `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// 商品ID类型。CUSTOM：外部编码；ITEM_SKU：淘系商品itemId_skuId组合形式。无SKU则为itemId_0
	IdType string `json:"id_type,omitempty" xml:"id_type,omitempty"`
}

var poolGoodsDetail = sync.Pool{
	New: func() any {
		return new(GoodsDetail)
	},
}

// GetGoodsDetail() 从对象池中获取GoodsDetail
func GetGoodsDetail() *GoodsDetail {
	return poolGoodsDetail.Get().(*GoodsDetail)
}

// ReleaseGoodsDetail 释放GoodsDetail
func ReleaseGoodsDetail(v *GoodsDetail) {
	v.GoodsId = ""
	v.Price = ""
	v.Quantity = ""
	v.IdType = ""
	poolGoodsDetail.Put(v)
}
