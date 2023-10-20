package security

import (
	"sync"
)

// OfficialAppApplyResponse 结构体
type OfficialAppApplyResponse struct {
	// message
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 错误码0-成功 400-参数错误 500-服务错误
	ErrCode int64 `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 提交状态 0- 提交成功 400-提交错误
	SubmitStatus int64 `json:"submit_status,omitempty" xml:"submit_status,omitempty"`
	// 请求是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolOfficialAppApplyResponse = sync.Pool{
	New: func() any {
		return new(OfficialAppApplyResponse)
	},
}

// GetOfficialAppApplyResponse() 从对象池中获取OfficialAppApplyResponse
func GetOfficialAppApplyResponse() *OfficialAppApplyResponse {
	return poolOfficialAppApplyResponse.Get().(*OfficialAppApplyResponse)
}

// ReleaseOfficialAppApplyResponse 释放OfficialAppApplyResponse
func ReleaseOfficialAppApplyResponse(v *OfficialAppApplyResponse) {
	v.ErrMsg = ""
	v.ErrCode = 0
	v.SubmitStatus = 0
	v.Success = false
	poolOfficialAppApplyResponse.Put(v)
}
