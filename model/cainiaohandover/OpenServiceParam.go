package cainiaohandover

import (
	"sync"
)

// OpenServiceParam 结构体
type OpenServiceParam struct {
	// DOOR_PICKUP:上门揽收；SELF_POST:自寄；SELF_SEND:自送；UNREACHABLE_RETURN:不可达退回；
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 不同物流服务的扩展信息
	Features *Features `json:"features,omitempty" xml:"features,omitempty"`
}

var poolOpenServiceParam = sync.Pool{
	New: func() any {
		return new(OpenServiceParam)
	},
}

// GetOpenServiceParam() 从对象池中获取OpenServiceParam
func GetOpenServiceParam() *OpenServiceParam {
	return poolOpenServiceParam.Get().(*OpenServiceParam)
}

// ReleaseOpenServiceParam 释放OpenServiceParam
func ReleaseOpenServiceParam(v *OpenServiceParam) {
	v.Code = ""
	v.Features = nil
	poolOpenServiceParam.Put(v)
}
