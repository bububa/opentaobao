package axintrade

import (
	"sync"
)

// AxinFundUpdateDto 结构体
type AxinFundUpdateDto struct {
	// 请求版本号，用于幂等校验
	ReqVersion string `json:"req_version,omitempty" xml:"req_version,omitempty"`
	// 扩展属性
	Attributes string `json:"attributes,omitempty" xml:"attributes,omitempty"`
	// 支付宝交易号
	AlipayOrderId string `json:"alipay_order_id,omitempty" xml:"alipay_order_id,omitempty"`
	// 订单ID
	OuterOrderId string `json:"outer_order_id,omitempty" xml:"outer_order_id,omitempty"`
	// 状态
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 资金单ID
	FundId int64 `json:"fund_id,omitempty" xml:"fund_id,omitempty"`
}

var poolAxinFundUpdateDto = sync.Pool{
	New: func() any {
		return new(AxinFundUpdateDto)
	},
}

// GetAxinFundUpdateDto() 从对象池中获取AxinFundUpdateDto
func GetAxinFundUpdateDto() *AxinFundUpdateDto {
	return poolAxinFundUpdateDto.Get().(*AxinFundUpdateDto)
}

// ReleaseAxinFundUpdateDto 释放AxinFundUpdateDto
func ReleaseAxinFundUpdateDto(v *AxinFundUpdateDto) {
	v.ReqVersion = ""
	v.Attributes = ""
	v.AlipayOrderId = ""
	v.OuterOrderId = ""
	v.Status = 0
	v.FundId = 0
	poolAxinFundUpdateDto.Put(v)
}
