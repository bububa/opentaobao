package fenxiao

import (
	"sync"
)

// TmallSupplychainChannelProductDownshelfResultDto 结构体
type TmallSupplychainChannelProductDownshelfResultDto struct {
	// 错误码
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 执行结果
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 下架结果
	Module bool `json:"module,omitempty" xml:"module,omitempty"`
}

var poolTmallSupplychainChannelProductDownshelfResultDto = sync.Pool{
	New: func() any {
		return new(TmallSupplychainChannelProductDownshelfResultDto)
	},
}

// GetTmallSupplychainChannelProductDownshelfResultDto() 从对象池中获取TmallSupplychainChannelProductDownshelfResultDto
func GetTmallSupplychainChannelProductDownshelfResultDto() *TmallSupplychainChannelProductDownshelfResultDto {
	return poolTmallSupplychainChannelProductDownshelfResultDto.Get().(*TmallSupplychainChannelProductDownshelfResultDto)
}

// ReleaseTmallSupplychainChannelProductDownshelfResultDto 释放TmallSupplychainChannelProductDownshelfResultDto
func ReleaseTmallSupplychainChannelProductDownshelfResultDto(v *TmallSupplychainChannelProductDownshelfResultDto) {
	v.ErrorMessage = ""
	v.ErrorCode = ""
	v.Success = false
	v.Module = false
	poolTmallSupplychainChannelProductDownshelfResultDto.Put(v)
}
