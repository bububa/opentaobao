package tmallnr

import (
	"sync"
)

// NrZqsGoodsReadyReqDto 结构体
type NrZqsGoodsReadyReqDto struct {
	// 配送人员姓名
	PerformerName string `json:"performer_name,omitempty" xml:"performer_name,omitempty"`
	// 配送人员电话号码
	PerformerPhone string `json:"performer_phone,omitempty" xml:"performer_phone,omitempty"`
	// 可选参数，编码
	TraceId string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
	// 提货码
	OuterGotCode string `json:"outer_got_code,omitempty" xml:"outer_got_code,omitempty"`
	// 经销商的姓名
	DealerName string `json:"dealer_name,omitempty" xml:"dealer_name,omitempty"`
	// 经销商的电话
	DealerPhone string `json:"dealer_phone,omitempty" xml:"dealer_phone,omitempty"`
	// 业务标识dds，zqs等
	BizIdentity string `json:"biz_identity,omitempty" xml:"biz_identity,omitempty"`
	// 卖家的sellerId
	SellerId int64 `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
	// 淘宝交易订单号
	BizOrderId int64 `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
}

var poolNrZqsGoodsReadyReqDto = sync.Pool{
	New: func() any {
		return new(NrZqsGoodsReadyReqDto)
	},
}

// GetNrZqsGoodsReadyReqDto() 从对象池中获取NrZqsGoodsReadyReqDto
func GetNrZqsGoodsReadyReqDto() *NrZqsGoodsReadyReqDto {
	return poolNrZqsGoodsReadyReqDto.Get().(*NrZqsGoodsReadyReqDto)
}

// ReleaseNrZqsGoodsReadyReqDto 释放NrZqsGoodsReadyReqDto
func ReleaseNrZqsGoodsReadyReqDto(v *NrZqsGoodsReadyReqDto) {
	v.PerformerName = ""
	v.PerformerPhone = ""
	v.TraceId = ""
	v.OuterGotCode = ""
	v.DealerName = ""
	v.DealerPhone = ""
	v.BizIdentity = ""
	v.SellerId = 0
	v.BizOrderId = 0
	poolNrZqsGoodsReadyReqDto.Put(v)
}
