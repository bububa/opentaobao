package alsc

import (
	"sync"
)

// BatchActiveCardOpenReq 结构体
type BatchActiveCardOpenReq struct {
	// 实体卡列表
	PhysicalCardIds []string `json:"physical_card_ids,omitempty" xml:"physical_card_ids>string,omitempty"`
	// 品牌id
	BrandId string `json:"brand_id,omitempty" xml:"brand_id,omitempty"`
	// 操作员id
	OperatorId string `json:"operator_id,omitempty" xml:"operator_id,omitempty"`
	// 外部品牌id,brandId与out_brand_id不可同时为空
	OutBrandId string `json:"out_brand_id,omitempty" xml:"out_brand_id,omitempty"`
	// 外部门店ID,shop_id和out_shop_id不可同时为空
	OutShopId string `json:"out_shop_id,omitempty" xml:"out_shop_id,omitempty"`
	// 请求id
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 门店id
	ShopId string `json:"shop_id,omitempty" xml:"shop_id,omitempty"`
}

var poolBatchActiveCardOpenReq = sync.Pool{
	New: func() any {
		return new(BatchActiveCardOpenReq)
	},
}

// GetBatchActiveCardOpenReq() 从对象池中获取BatchActiveCardOpenReq
func GetBatchActiveCardOpenReq() *BatchActiveCardOpenReq {
	return poolBatchActiveCardOpenReq.Get().(*BatchActiveCardOpenReq)
}

// ReleaseBatchActiveCardOpenReq 释放BatchActiveCardOpenReq
func ReleaseBatchActiveCardOpenReq(v *BatchActiveCardOpenReq) {
	v.PhysicalCardIds = v.PhysicalCardIds[:0]
	v.BrandId = ""
	v.OperatorId = ""
	v.OutBrandId = ""
	v.OutShopId = ""
	v.RequestId = ""
	v.ShopId = ""
	poolBatchActiveCardOpenReq.Put(v)
}
