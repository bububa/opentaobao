package fenxiao

import (
	"sync"
)

// TmallSupplychainChannelProductQuantityUpdateResultDto 结构体
type TmallSupplychainChannelProductQuantityUpdateResultDto struct {
	// 错误码
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 更新内容
	Module *TopProductQuantityResult `json:"module,omitempty" xml:"module,omitempty"`
	// 执行结果
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTmallSupplychainChannelProductQuantityUpdateResultDto = sync.Pool{
	New: func() any {
		return new(TmallSupplychainChannelProductQuantityUpdateResultDto)
	},
}

// GetTmallSupplychainChannelProductQuantityUpdateResultDto() 从对象池中获取TmallSupplychainChannelProductQuantityUpdateResultDto
func GetTmallSupplychainChannelProductQuantityUpdateResultDto() *TmallSupplychainChannelProductQuantityUpdateResultDto {
	return poolTmallSupplychainChannelProductQuantityUpdateResultDto.Get().(*TmallSupplychainChannelProductQuantityUpdateResultDto)
}

// ReleaseTmallSupplychainChannelProductQuantityUpdateResultDto 释放TmallSupplychainChannelProductQuantityUpdateResultDto
func ReleaseTmallSupplychainChannelProductQuantityUpdateResultDto(v *TmallSupplychainChannelProductQuantityUpdateResultDto) {
	v.ErrorMessage = ""
	v.ErrorCode = ""
	v.Module = nil
	v.Success = false
	poolTmallSupplychainChannelProductQuantityUpdateResultDto.Put(v)
}
