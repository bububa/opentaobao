package icburfq

import (
	"sync"
)

// RfqQuotationPriceRemoteDto 结构体
type RfqQuotationPriceRemoteDto struct {
	// 备注
	Remark string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 样品运费支付方
	Payment string `json:"payment,omitempty" xml:"payment,omitempty"`
	// 是否是免费
	IsFree string `json:"is_free,omitempty" xml:"is_free,omitempty"`
	// 是否提供样本
	IsSupport string `json:"is_support,omitempty" xml:"is_support,omitempty"`
	// 预计时间
	EstimatedDate float64 `json:"estimated_date,omitempty" xml:"estimated_date,omitempty"`
}

var poolRfqQuotationPriceRemoteDto = sync.Pool{
	New: func() any {
		return new(RfqQuotationPriceRemoteDto)
	},
}

// GetRfqQuotationPriceRemoteDto() 从对象池中获取RfqQuotationPriceRemoteDto
func GetRfqQuotationPriceRemoteDto() *RfqQuotationPriceRemoteDto {
	return poolRfqQuotationPriceRemoteDto.Get().(*RfqQuotationPriceRemoteDto)
}

// ReleaseRfqQuotationPriceRemoteDto 释放RfqQuotationPriceRemoteDto
func ReleaseRfqQuotationPriceRemoteDto(v *RfqQuotationPriceRemoteDto) {
	v.Remark = ""
	v.Payment = ""
	v.IsFree = ""
	v.IsSupport = ""
	v.EstimatedDate = 0
	poolRfqQuotationPriceRemoteDto.Put(v)
}
