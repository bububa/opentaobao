package mos

import (
	"sync"
)

// GoodsOutboundDto 结构体
type GoodsOutboundDto struct {
	// 出库明细
	OutboundDetails []OutboundDetailDto `json:"outbound_details,omitempty" xml:"outbound_details>outbound_detail_dto,omitempty"`
	// 配送员
	DelivererName string `json:"deliverer_name,omitempty" xml:"deliverer_name,omitempty"`
	// 配送员电话
	DelivererPhone string `json:"deliverer_phone,omitempty" xml:"deliverer_phone,omitempty"`
	// 物流公司代码
	LogisticsCompanyCode string `json:"logistics_company_code,omitempty" xml:"logistics_company_code,omitempty"`
	// 物流单号
	LogisticsNo string `json:"logistics_no,omitempty" xml:"logistics_no,omitempty"`
	// 发货时间
	SendOutTime string `json:"send_out_time,omitempty" xml:"send_out_time,omitempty"`
	// OC订单号
	TradeNo string `json:"trade_no,omitempty" xml:"trade_no,omitempty"`
}

var poolGoodsOutboundDto = sync.Pool{
	New: func() any {
		return new(GoodsOutboundDto)
	},
}

// GetGoodsOutboundDto() 从对象池中获取GoodsOutboundDto
func GetGoodsOutboundDto() *GoodsOutboundDto {
	return poolGoodsOutboundDto.Get().(*GoodsOutboundDto)
}

// ReleaseGoodsOutboundDto 释放GoodsOutboundDto
func ReleaseGoodsOutboundDto(v *GoodsOutboundDto) {
	v.OutboundDetails = v.OutboundDetails[:0]
	v.DelivererName = ""
	v.DelivererPhone = ""
	v.LogisticsCompanyCode = ""
	v.LogisticsNo = ""
	v.SendOutTime = ""
	v.TradeNo = ""
	poolGoodsOutboundDto.Put(v)
}
