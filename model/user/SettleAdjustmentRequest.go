package user

import (
	"sync"
)

// SettleAdjustmentRequest 结构体
type SettleAdjustmentRequest struct {
	// 计价因子，填写规则为：1、有计价因子场景：{name:计价因子名称 ,value:数量｝如示例。2、没有计价因子场景：填默认值：｛name:计价因子,value:0｝
	PriceFactors []SettlementPriceFactor `json:"price_factors,omitempty" xml:"price_factors>settlement_price_factor,omitempty"`
	// 调整原因描述
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	// 调整原因图片url,最后不用加分号，最多三条
	PictureUrls string `json:"picture_urls,omitempty" xml:"picture_urls,omitempty"`
	// 调整费用，单位分
	Cost int64 `json:"cost,omitempty" xml:"cost,omitempty"`
	// 调整单ID
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 调整单类型：1,配件费;2,不符单费;3,拆旧费;4,二次上门;5,胶费;6,打孔费;7,层高费;8,远程费;9,单外费;10,其他
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
}

var poolSettleAdjustmentRequest = sync.Pool{
	New: func() any {
		return new(SettleAdjustmentRequest)
	},
}

// GetSettleAdjustmentRequest() 从对象池中获取SettleAdjustmentRequest
func GetSettleAdjustmentRequest() *SettleAdjustmentRequest {
	return poolSettleAdjustmentRequest.Get().(*SettleAdjustmentRequest)
}

// ReleaseSettleAdjustmentRequest 释放SettleAdjustmentRequest
func ReleaseSettleAdjustmentRequest(v *SettleAdjustmentRequest) {
	v.PriceFactors = v.PriceFactors[:0]
	v.Description = ""
	v.PictureUrls = ""
	v.Cost = 0
	v.Id = 0
	v.Type = 0
	poolSettleAdjustmentRequest.Put(v)
}
