package perfect

import (
	"sync"
)

// BoxCodeResponse 结构体
type BoxCodeResponse struct {
	// 1
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 1
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 箱号
	BoxCode string `json:"box_code,omitempty" xml:"box_code,omitempty"`
	// 1
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolBoxCodeResponse = sync.Pool{
	New: func() any {
		return new(BoxCodeResponse)
	},
}

// GetBoxCodeResponse() 从对象池中获取BoxCodeResponse
func GetBoxCodeResponse() *BoxCodeResponse {
	return poolBoxCodeResponse.Get().(*BoxCodeResponse)
}

// ReleaseBoxCodeResponse 释放BoxCodeResponse
func ReleaseBoxCodeResponse(v *BoxCodeResponse) {
	v.ErrorCode = ""
	v.ErrorMessage = ""
	v.BoxCode = ""
	v.Success = false
	poolBoxCodeResponse.Put(v)
}
