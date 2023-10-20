package fenxiao

import (
	"sync"
)

// TmallSupplychainChannelProductReleaseResultDto 结构体
type TmallSupplychainChannelProductReleaseResultDto struct {
	// 链路ID
	TraceId string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
	// 异常名
	ExpName string `json:"exp_name,omitempty" xml:"exp_name,omitempty"`
	// 重定向url
	RedirectUrl string `json:"redirect_url,omitempty" xml:"redirect_url,omitempty"`
	// 错误码
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 执行结果
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTmallSupplychainChannelProductReleaseResultDto = sync.Pool{
	New: func() any {
		return new(TmallSupplychainChannelProductReleaseResultDto)
	},
}

// GetTmallSupplychainChannelProductReleaseResultDto() 从对象池中获取TmallSupplychainChannelProductReleaseResultDto
func GetTmallSupplychainChannelProductReleaseResultDto() *TmallSupplychainChannelProductReleaseResultDto {
	return poolTmallSupplychainChannelProductReleaseResultDto.Get().(*TmallSupplychainChannelProductReleaseResultDto)
}

// ReleaseTmallSupplychainChannelProductReleaseResultDto 释放TmallSupplychainChannelProductReleaseResultDto
func ReleaseTmallSupplychainChannelProductReleaseResultDto(v *TmallSupplychainChannelProductReleaseResultDto) {
	v.TraceId = ""
	v.ExpName = ""
	v.RedirectUrl = ""
	v.ErrorMessage = ""
	v.ErrorCode = ""
	v.Success = false
	poolTmallSupplychainChannelProductReleaseResultDto.Put(v)
}
