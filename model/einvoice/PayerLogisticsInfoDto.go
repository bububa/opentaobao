package einvoice

import (
	"sync"
)

// PayerLogisticsInfoDto 结构体
type PayerLogisticsInfoDto struct {
	// 收件人地址
	ContactAddr string `json:"contact_addr,omitempty" xml:"contact_addr,omitempty"`
	// 收件人电话
	ContactMobile string `json:"contact_mobile,omitempty" xml:"contact_mobile,omitempty"`
	// 收件人姓名
	ContactName string `json:"contact_name,omitempty" xml:"contact_name,omitempty"`
}

var poolPayerLogisticsInfoDto = sync.Pool{
	New: func() any {
		return new(PayerLogisticsInfoDto)
	},
}

// GetPayerLogisticsInfoDto() 从对象池中获取PayerLogisticsInfoDto
func GetPayerLogisticsInfoDto() *PayerLogisticsInfoDto {
	return poolPayerLogisticsInfoDto.Get().(*PayerLogisticsInfoDto)
}

// ReleasePayerLogisticsInfoDto 释放PayerLogisticsInfoDto
func ReleasePayerLogisticsInfoDto(v *PayerLogisticsInfoDto) {
	v.ContactAddr = ""
	v.ContactMobile = ""
	v.ContactName = ""
	poolPayerLogisticsInfoDto.Put(v)
}
