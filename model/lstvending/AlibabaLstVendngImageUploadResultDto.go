package lstvending

import (
	"sync"
)

// AlibabaLstVendngImageUploadResultDto 结构体
type AlibabaLstVendngImageUploadResultDto struct {
	// 错误信息
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 图片上传信息
	Module *VendingImageDto `json:"module,omitempty" xml:"module,omitempty"`
	// 是否处理成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaLstVendngImageUploadResultDto = sync.Pool{
	New: func() any {
		return new(AlibabaLstVendngImageUploadResultDto)
	},
}

// GetAlibabaLstVendngImageUploadResultDto() 从对象池中获取AlibabaLstVendngImageUploadResultDto
func GetAlibabaLstVendngImageUploadResultDto() *AlibabaLstVendngImageUploadResultDto {
	return poolAlibabaLstVendngImageUploadResultDto.Get().(*AlibabaLstVendngImageUploadResultDto)
}

// ReleaseAlibabaLstVendngImageUploadResultDto 释放AlibabaLstVendngImageUploadResultDto
func ReleaseAlibabaLstVendngImageUploadResultDto(v *AlibabaLstVendngImageUploadResultDto) {
	v.ErrorMessage = ""
	v.ErrorCode = ""
	v.Module = nil
	v.Success = false
	poolAlibabaLstVendngImageUploadResultDto.Put(v)
}
