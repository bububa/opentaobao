package alitripdivisions

import (
	"sync"
)

// AlitripPlatformDivisionsQuerybyparentidResult 结构体
type AlitripPlatformDivisionsQuerybyparentidResult struct {
	// mapping code
	MappingCode string `json:"mapping_code,omitempty" xml:"mapping_code,omitempty"`
	// msgCode
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// msgInfo
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// http status code
	HttpStatusCode int64 `json:"http_status_code,omitempty" xml:"http_status_code,omitempty"`
	// model
	Model *TrdiDivisionBasicListVo `json:"model,omitempty" xml:"model,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlitripPlatformDivisionsQuerybyparentidResult = sync.Pool{
	New: func() any {
		return new(AlitripPlatformDivisionsQuerybyparentidResult)
	},
}

// GetAlitripPlatformDivisionsQuerybyparentidResult() 从对象池中获取AlitripPlatformDivisionsQuerybyparentidResult
func GetAlitripPlatformDivisionsQuerybyparentidResult() *AlitripPlatformDivisionsQuerybyparentidResult {
	return poolAlitripPlatformDivisionsQuerybyparentidResult.Get().(*AlitripPlatformDivisionsQuerybyparentidResult)
}

// ReleaseAlitripPlatformDivisionsQuerybyparentidResult 释放AlitripPlatformDivisionsQuerybyparentidResult
func ReleaseAlitripPlatformDivisionsQuerybyparentidResult(v *AlitripPlatformDivisionsQuerybyparentidResult) {
	v.MappingCode = ""
	v.MsgCode = ""
	v.MsgInfo = ""
	v.HttpStatusCode = 0
	v.Model = nil
	v.Success = false
	poolAlitripPlatformDivisionsQuerybyparentidResult.Put(v)
}
