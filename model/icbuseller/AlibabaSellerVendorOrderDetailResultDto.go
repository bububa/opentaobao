package icbuseller

import (
	"sync"
)

// AlibabaSellerVendorOrderDetailResultDto 结构体
type AlibabaSellerVendorOrderDetailResultDto struct {
	// 接口返回对象
	Dto *OpenTradeDetailDto `json:"dto,omitempty" xml:"dto,omitempty"`
}

var poolAlibabaSellerVendorOrderDetailResultDto = sync.Pool{
	New: func() any {
		return new(AlibabaSellerVendorOrderDetailResultDto)
	},
}

// GetAlibabaSellerVendorOrderDetailResultDto() 从对象池中获取AlibabaSellerVendorOrderDetailResultDto
func GetAlibabaSellerVendorOrderDetailResultDto() *AlibabaSellerVendorOrderDetailResultDto {
	return poolAlibabaSellerVendorOrderDetailResultDto.Get().(*AlibabaSellerVendorOrderDetailResultDto)
}

// ReleaseAlibabaSellerVendorOrderDetailResultDto 释放AlibabaSellerVendorOrderDetailResultDto
func ReleaseAlibabaSellerVendorOrderDetailResultDto(v *AlibabaSellerVendorOrderDetailResultDto) {
	v.Dto = nil
	poolAlibabaSellerVendorOrderDetailResultDto.Put(v)
}
