package fenxiao

import (
	"sync"
)

// TmallSupplychainChannelProductPriceGetResultDto 结构体
type TmallSupplychainChannelProductPriceGetResultDto struct {
	// 错误码
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 询价结果
	Module *TopProductPriceResult `json:"module,omitempty" xml:"module,omitempty"`
	// 执行结果
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTmallSupplychainChannelProductPriceGetResultDto = sync.Pool{
	New: func() any {
		return new(TmallSupplychainChannelProductPriceGetResultDto)
	},
}

// GetTmallSupplychainChannelProductPriceGetResultDto() 从对象池中获取TmallSupplychainChannelProductPriceGetResultDto
func GetTmallSupplychainChannelProductPriceGetResultDto() *TmallSupplychainChannelProductPriceGetResultDto {
	return poolTmallSupplychainChannelProductPriceGetResultDto.Get().(*TmallSupplychainChannelProductPriceGetResultDto)
}

// ReleaseTmallSupplychainChannelProductPriceGetResultDto 释放TmallSupplychainChannelProductPriceGetResultDto
func ReleaseTmallSupplychainChannelProductPriceGetResultDto(v *TmallSupplychainChannelProductPriceGetResultDto) {
	v.ErrorMessage = ""
	v.ErrorCode = ""
	v.Module = nil
	v.Success = false
	poolTmallSupplychainChannelProductPriceGetResultDto.Put(v)
}
