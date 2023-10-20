package btrip

import (
	"sync"
)

// OpenApiZzdSearchRq 结构体
type OpenApiZzdSearchRq struct {
	// 第三方企业ID
	ThirdpartCorpId string `json:"thirdpart_corp_id,omitempty" xml:"thirdpart_corp_id,omitempty"`
	// 结算结束日期
	EndDate string `json:"end_date,omitempty" xml:"end_date,omitempty"`
	// 第三方用户ID
	UserId string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// 第三方交易号
	TradeId string `json:"trade_id,omitempty" xml:"trade_id,omitempty"`
	// 结算开始时间
	StartDate string `json:"start_date,omitempty" xml:"start_date,omitempty"`
	// 订单ID
	OrderId int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 每页数量
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 商旅申请单号
	ApplyId int64 `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
	// 第几页
	Page int64 `json:"page,omitempty" xml:"page,omitempty"`
}

var poolOpenApiZzdSearchRq = sync.Pool{
	New: func() any {
		return new(OpenApiZzdSearchRq)
	},
}

// GetOpenApiZzdSearchRq() 从对象池中获取OpenApiZzdSearchRq
func GetOpenApiZzdSearchRq() *OpenApiZzdSearchRq {
	return poolOpenApiZzdSearchRq.Get().(*OpenApiZzdSearchRq)
}

// ReleaseOpenApiZzdSearchRq 释放OpenApiZzdSearchRq
func ReleaseOpenApiZzdSearchRq(v *OpenApiZzdSearchRq) {
	v.ThirdpartCorpId = ""
	v.EndDate = ""
	v.UserId = ""
	v.TradeId = ""
	v.StartDate = ""
	v.OrderId = 0
	v.PageSize = 0
	v.ApplyId = 0
	v.Page = 0
	poolOpenApiZzdSearchRq.Put(v)
}
