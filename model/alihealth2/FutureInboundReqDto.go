package alihealth2

import (
	"sync"
)

// FutureInboundReqDto 结构体
type FutureInboundReqDto struct {
	// 请求明细,数量需与期货采购一致
	Items []FutureInboundItem `json:"items,omitempty" xml:"items>future_inbound_item,omitempty"`
	// 期货采购ID
	FuturePurchaseId string `json:"future_purchase_id,omitempty" xml:"future_purchase_id,omitempty"`
	// 请求流水号, 长度不能超过64, 相同的请求流水会被幂等
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 仓库编码
	StoreCode string `json:"store_code,omitempty" xml:"store_code,omitempty"`
	// 供应商ID
	SupplierId int64 `json:"supplier_id,omitempty" xml:"supplier_id,omitempty"`
	// 是否最后一次入库
	LastInbound bool `json:"last_inbound,omitempty" xml:"last_inbound,omitempty"`
}

var poolFutureInboundReqDto = sync.Pool{
	New: func() any {
		return new(FutureInboundReqDto)
	},
}

// GetFutureInboundReqDto() 从对象池中获取FutureInboundReqDto
func GetFutureInboundReqDto() *FutureInboundReqDto {
	return poolFutureInboundReqDto.Get().(*FutureInboundReqDto)
}

// ReleaseFutureInboundReqDto 释放FutureInboundReqDto
func ReleaseFutureInboundReqDto(v *FutureInboundReqDto) {
	v.Items = v.Items[:0]
	v.FuturePurchaseId = ""
	v.RequestId = ""
	v.StoreCode = ""
	v.SupplierId = 0
	v.LastInbound = false
	poolFutureInboundReqDto.Put(v)
}
