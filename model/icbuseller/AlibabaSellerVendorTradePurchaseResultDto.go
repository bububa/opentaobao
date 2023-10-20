package icbuseller

import (
	"sync"
)

// AlibabaSellerVendorTradePurchaseResultDto 结构体
type AlibabaSellerVendorTradePurchaseResultDto struct {
	// 授权订单集合
	Dtos []TradePurchaseDto `json:"dtos,omitempty" xml:"dtos>trade_purchase_dto,omitempty"`
	// 描述
	ExecDescription string `json:"exec_description,omitempty" xml:"exec_description,omitempty"`
	// 接口状态
	ReturnCode int64 `json:"return_code,omitempty" xml:"return_code,omitempty"`
	// 执行结果
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaSellerVendorTradePurchaseResultDto = sync.Pool{
	New: func() any {
		return new(AlibabaSellerVendorTradePurchaseResultDto)
	},
}

// GetAlibabaSellerVendorTradePurchaseResultDto() 从对象池中获取AlibabaSellerVendorTradePurchaseResultDto
func GetAlibabaSellerVendorTradePurchaseResultDto() *AlibabaSellerVendorTradePurchaseResultDto {
	return poolAlibabaSellerVendorTradePurchaseResultDto.Get().(*AlibabaSellerVendorTradePurchaseResultDto)
}

// ReleaseAlibabaSellerVendorTradePurchaseResultDto 释放AlibabaSellerVendorTradePurchaseResultDto
func ReleaseAlibabaSellerVendorTradePurchaseResultDto(v *AlibabaSellerVendorTradePurchaseResultDto) {
	v.Dtos = v.Dtos[:0]
	v.ExecDescription = ""
	v.ReturnCode = 0
	v.Success = false
	poolAlibabaSellerVendorTradePurchaseResultDto.Put(v)
}
