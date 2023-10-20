package ascpchannel

import (
	"sync"
)

// AlibabaAscpChannelSupplierProductDetailResultDto 结构体
type AlibabaAscpChannelSupplierProductDetailResultDto struct {
	// 错误码
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 返回详情
	Module *ProductDetailQueryResponseForSupplier `json:"module,omitempty" xml:"module,omitempty"`
	// 执行结果
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAscpChannelSupplierProductDetailResultDto = sync.Pool{
	New: func() any {
		return new(AlibabaAscpChannelSupplierProductDetailResultDto)
	},
}

// GetAlibabaAscpChannelSupplierProductDetailResultDto() 从对象池中获取AlibabaAscpChannelSupplierProductDetailResultDto
func GetAlibabaAscpChannelSupplierProductDetailResultDto() *AlibabaAscpChannelSupplierProductDetailResultDto {
	return poolAlibabaAscpChannelSupplierProductDetailResultDto.Get().(*AlibabaAscpChannelSupplierProductDetailResultDto)
}

// ReleaseAlibabaAscpChannelSupplierProductDetailResultDto 释放AlibabaAscpChannelSupplierProductDetailResultDto
func ReleaseAlibabaAscpChannelSupplierProductDetailResultDto(v *AlibabaAscpChannelSupplierProductDetailResultDto) {
	v.ErrorMessage = ""
	v.ErrorCode = ""
	v.Module = nil
	v.Success = false
	poolAlibabaAscpChannelSupplierProductDetailResultDto.Put(v)
}
