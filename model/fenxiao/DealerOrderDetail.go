package fenxiao

import (
	"sync"
)

// DealerOrderDetail 结构体
type DealerOrderDetail struct {
	// 属性信息列表，key-value形式。如attr_key为storeCode，attr_value则为仓库编码。
	Features []Feature `json:"features,omitempty" xml:"features>feature,omitempty"`
	// 最终价格(精确到2位小数;单位:元。如:200.07，表示:200元7分 )
	FinalPrice string `json:"final_price,omitempty" xml:"final_price,omitempty"`
	// 产品规格
	SkuSpec string `json:"sku_spec,omitempty" xml:"sku_spec,omitempty"`
	// 原始价格(精确到2位小数;单位:元。如:200.07，表示:200元7分 )
	OriginalPrice string `json:"original_price,omitempty" xml:"original_price,omitempty"`
	// 产品快照url
	SnapshotUrl string `json:"snapshot_url,omitempty" xml:"snapshot_url,omitempty"`
	// 产品标题
	ProductTitle string `json:"product_title,omitempty" xml:"product_title,omitempty"`
	// 金额小计(采购数量乘以最终价格。精确到2位小数;单位:元。如:200.07，表示:200元7分 )
	PriceCount string `json:"price_count,omitempty" xml:"price_count,omitempty"`
	// 商家编码，是sku则为sku的商家编码，否则是产品的商家编码
	SkuNumber string `json:"sku_number,omitempty" xml:"sku_number,omitempty"`
	// 产品id
	ProductId int64 `json:"product_id,omitempty" xml:"product_id,omitempty"`
	// 经销采购单明细id
	DealerDetailId int64 `json:"dealer_detail_id,omitempty" xml:"dealer_detail_id,omitempty"`
	// 经销采购单编号
	DealerOrderId int64 `json:"dealer_order_id,omitempty" xml:"dealer_order_id,omitempty"`
	// 采购数量
	Quantity int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// sku id
	SkuId int64 `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	// 该条明细是否已删除。true：已删除；false：未删除。
	IsDeleted bool `json:"is_deleted,omitempty" xml:"is_deleted,omitempty"`
}

var poolDealerOrderDetail = sync.Pool{
	New: func() any {
		return new(DealerOrderDetail)
	},
}

// GetDealerOrderDetail() 从对象池中获取DealerOrderDetail
func GetDealerOrderDetail() *DealerOrderDetail {
	return poolDealerOrderDetail.Get().(*DealerOrderDetail)
}

// ReleaseDealerOrderDetail 释放DealerOrderDetail
func ReleaseDealerOrderDetail(v *DealerOrderDetail) {
	v.Features = v.Features[:0]
	v.FinalPrice = ""
	v.SkuSpec = ""
	v.OriginalPrice = ""
	v.SnapshotUrl = ""
	v.ProductTitle = ""
	v.PriceCount = ""
	v.SkuNumber = ""
	v.ProductId = 0
	v.DealerDetailId = 0
	v.DealerOrderId = 0
	v.Quantity = 0
	v.SkuId = 0
	v.IsDeleted = false
	poolDealerOrderDetail.Put(v)
}
