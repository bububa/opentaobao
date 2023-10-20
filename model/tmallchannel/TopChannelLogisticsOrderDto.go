package tmallchannel

import (
	"sync"
)

// TopChannelLogisticsOrderDto 结构体
type TopChannelLogisticsOrderDto struct {
	// 物流单号
	LogisticsOrderNo string `json:"logistics_order_no,omitempty" xml:"logistics_order_no,omitempty"`
	// 物流公司编号
	LogisticsCompanyCode string `json:"logistics_company_code,omitempty" xml:"logistics_company_code,omitempty"`
	// 物流公司名称
	LogisticsCompanyName string `json:"logistics_company_name,omitempty" xml:"logistics_company_name,omitempty"`
}

var poolTopChannelLogisticsOrderDto = sync.Pool{
	New: func() any {
		return new(TopChannelLogisticsOrderDto)
	},
}

// GetTopChannelLogisticsOrderDto() 从对象池中获取TopChannelLogisticsOrderDto
func GetTopChannelLogisticsOrderDto() *TopChannelLogisticsOrderDto {
	return poolTopChannelLogisticsOrderDto.Get().(*TopChannelLogisticsOrderDto)
}

// ReleaseTopChannelLogisticsOrderDto 释放TopChannelLogisticsOrderDto
func ReleaseTopChannelLogisticsOrderDto(v *TopChannelLogisticsOrderDto) {
	v.LogisticsOrderNo = ""
	v.LogisticsCompanyCode = ""
	v.LogisticsCompanyName = ""
	poolTopChannelLogisticsOrderDto.Put(v)
}
