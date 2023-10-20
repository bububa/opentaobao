package alitripmerchant

import (
	"sync"
)

// BindingMerchantIdUserDto 结构体
type BindingMerchantIdUserDto struct {
	// 1
	MerchantId string `json:"merchant_id,omitempty" xml:"merchant_id,omitempty"`
	// 1
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 1
	Type string `json:"type,omitempty" xml:"type,omitempty"`
}

var poolBindingMerchantIdUserDto = sync.Pool{
	New: func() any {
		return new(BindingMerchantIdUserDto)
	},
}

// GetBindingMerchantIdUserDto() 从对象池中获取BindingMerchantIdUserDto
func GetBindingMerchantIdUserDto() *BindingMerchantIdUserDto {
	return poolBindingMerchantIdUserDto.Get().(*BindingMerchantIdUserDto)
}

// ReleaseBindingMerchantIdUserDto 释放BindingMerchantIdUserDto
func ReleaseBindingMerchantIdUserDto(v *BindingMerchantIdUserDto) {
	v.MerchantId = ""
	v.Code = ""
	v.Type = ""
	poolBindingMerchantIdUserDto.Put(v)
}
