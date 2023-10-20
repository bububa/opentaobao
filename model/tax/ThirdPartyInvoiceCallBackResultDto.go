package tax

import (
	"sync"
)

// ThirdPartyInvoiceCallBackResultDto 结构体
type ThirdPartyInvoiceCallBackResultDto struct {
	// 具体明细列表
	ValueList []ResultItem `json:"value_list,omitempty" xml:"value_list>result_item,omitempty"`
}

var poolThirdPartyInvoiceCallBackResultDto = sync.Pool{
	New: func() any {
		return new(ThirdPartyInvoiceCallBackResultDto)
	},
}

// GetThirdPartyInvoiceCallBackResultDto() 从对象池中获取ThirdPartyInvoiceCallBackResultDto
func GetThirdPartyInvoiceCallBackResultDto() *ThirdPartyInvoiceCallBackResultDto {
	return poolThirdPartyInvoiceCallBackResultDto.Get().(*ThirdPartyInvoiceCallBackResultDto)
}

// ReleaseThirdPartyInvoiceCallBackResultDto 释放ThirdPartyInvoiceCallBackResultDto
func ReleaseThirdPartyInvoiceCallBackResultDto(v *ThirdPartyInvoiceCallBackResultDto) {
	v.ValueList = v.ValueList[:0]
	poolThirdPartyInvoiceCallBackResultDto.Put(v)
}
