package ascpqcc

import (
	"sync"
)

// UpdateSampleResponse 结构体
type UpdateSampleResponse struct {
	// 业务系统错误信息
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 业务系统错误编号
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 业务系统是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolUpdateSampleResponse = sync.Pool{
	New: func() any {
		return new(UpdateSampleResponse)
	},
}

// GetUpdateSampleResponse() 从对象池中获取UpdateSampleResponse
func GetUpdateSampleResponse() *UpdateSampleResponse {
	return poolUpdateSampleResponse.Get().(*UpdateSampleResponse)
}

// ReleaseUpdateSampleResponse 释放UpdateSampleResponse
func ReleaseUpdateSampleResponse(v *UpdateSampleResponse) {
	v.ErrorMessage = ""
	v.ErrorCode = ""
	v.Success = false
	poolUpdateSampleResponse.Put(v)
}
