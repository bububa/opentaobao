package fenxiao

import (
	"sync"
)

// TmallSupplychainChannelProductQuantityGetResultDto 结构体
type TmallSupplychainChannelProductQuantityGetResultDto struct {
	// 库存数量
	Module string `json:"module,omitempty" xml:"module,omitempty"`
	// 错误码
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 执行结果
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTmallSupplychainChannelProductQuantityGetResultDto = sync.Pool{
	New: func() any {
		return new(TmallSupplychainChannelProductQuantityGetResultDto)
	},
}

// GetTmallSupplychainChannelProductQuantityGetResultDto() 从对象池中获取TmallSupplychainChannelProductQuantityGetResultDto
func GetTmallSupplychainChannelProductQuantityGetResultDto() *TmallSupplychainChannelProductQuantityGetResultDto {
	return poolTmallSupplychainChannelProductQuantityGetResultDto.Get().(*TmallSupplychainChannelProductQuantityGetResultDto)
}

// ReleaseTmallSupplychainChannelProductQuantityGetResultDto 释放TmallSupplychainChannelProductQuantityGetResultDto
func ReleaseTmallSupplychainChannelProductQuantityGetResultDto(v *TmallSupplychainChannelProductQuantityGetResultDto) {
	v.Module = ""
	v.ErrorMessage = ""
	v.ErrorCode = ""
	v.Success = false
	poolTmallSupplychainChannelProductQuantityGetResultDto.Put(v)
}
