package fenxiao

import (
	"sync"
)

// TmallSupplychainChannelProductReleaseStatusGetResultDto 结构体
type TmallSupplychainChannelProductReleaseStatusGetResultDto struct {
	// 错误码
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 查询结果
	Module *TopProductStatusResult `json:"module,omitempty" xml:"module,omitempty"`
	// 执行结果
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTmallSupplychainChannelProductReleaseStatusGetResultDto = sync.Pool{
	New: func() any {
		return new(TmallSupplychainChannelProductReleaseStatusGetResultDto)
	},
}

// GetTmallSupplychainChannelProductReleaseStatusGetResultDto() 从对象池中获取TmallSupplychainChannelProductReleaseStatusGetResultDto
func GetTmallSupplychainChannelProductReleaseStatusGetResultDto() *TmallSupplychainChannelProductReleaseStatusGetResultDto {
	return poolTmallSupplychainChannelProductReleaseStatusGetResultDto.Get().(*TmallSupplychainChannelProductReleaseStatusGetResultDto)
}

// ReleaseTmallSupplychainChannelProductReleaseStatusGetResultDto 释放TmallSupplychainChannelProductReleaseStatusGetResultDto
func ReleaseTmallSupplychainChannelProductReleaseStatusGetResultDto(v *TmallSupplychainChannelProductReleaseStatusGetResultDto) {
	v.ErrorMessage = ""
	v.ErrorCode = ""
	v.Module = nil
	v.Success = false
	poolTmallSupplychainChannelProductReleaseStatusGetResultDto.Put(v)
}
