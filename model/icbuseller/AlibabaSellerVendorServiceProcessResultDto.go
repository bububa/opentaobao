package icbuseller

import (
	"sync"
)

// AlibabaSellerVendorServiceProcessResultDto 结构体
type AlibabaSellerVendorServiceProcessResultDto struct {
	// 返回集合
	List []VendorMerchantRecordBaseDto `json:"list,omitempty" xml:"list>vendor_merchant_record_base_dto,omitempty"`
	// 异常说明
	ExecDescription string `json:"exec_description,omitempty" xml:"exec_description,omitempty"`
	// 状态码
	ReturnCode int64 `json:"return_code,omitempty" xml:"return_code,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaSellerVendorServiceProcessResultDto = sync.Pool{
	New: func() any {
		return new(AlibabaSellerVendorServiceProcessResultDto)
	},
}

// GetAlibabaSellerVendorServiceProcessResultDto() 从对象池中获取AlibabaSellerVendorServiceProcessResultDto
func GetAlibabaSellerVendorServiceProcessResultDto() *AlibabaSellerVendorServiceProcessResultDto {
	return poolAlibabaSellerVendorServiceProcessResultDto.Get().(*AlibabaSellerVendorServiceProcessResultDto)
}

// ReleaseAlibabaSellerVendorServiceProcessResultDto 释放AlibabaSellerVendorServiceProcessResultDto
func ReleaseAlibabaSellerVendorServiceProcessResultDto(v *AlibabaSellerVendorServiceProcessResultDto) {
	v.List = v.List[:0]
	v.ExecDescription = ""
	v.ReturnCode = 0
	v.Success = false
	poolAlibabaSellerVendorServiceProcessResultDto.Put(v)
}
