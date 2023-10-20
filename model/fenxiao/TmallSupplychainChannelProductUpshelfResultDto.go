package fenxiao

import (
	"sync"
)

// TmallSupplychainChannelProductUpshelfResultDto 结构体
type TmallSupplychainChannelProductUpshelfResultDto struct {
	// 错误码
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 执行结果
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 上架结果
	Module bool `json:"module,omitempty" xml:"module,omitempty"`
}

var poolTmallSupplychainChannelProductUpshelfResultDto = sync.Pool{
	New: func() any {
		return new(TmallSupplychainChannelProductUpshelfResultDto)
	},
}

// GetTmallSupplychainChannelProductUpshelfResultDto() 从对象池中获取TmallSupplychainChannelProductUpshelfResultDto
func GetTmallSupplychainChannelProductUpshelfResultDto() *TmallSupplychainChannelProductUpshelfResultDto {
	return poolTmallSupplychainChannelProductUpshelfResultDto.Get().(*TmallSupplychainChannelProductUpshelfResultDto)
}

// ReleaseTmallSupplychainChannelProductUpshelfResultDto 释放TmallSupplychainChannelProductUpshelfResultDto
func ReleaseTmallSupplychainChannelProductUpshelfResultDto(v *TmallSupplychainChannelProductUpshelfResultDto) {
	v.ErrorMessage = ""
	v.ErrorCode = ""
	v.Success = false
	v.Module = false
	poolTmallSupplychainChannelProductUpshelfResultDto.Put(v)
}
