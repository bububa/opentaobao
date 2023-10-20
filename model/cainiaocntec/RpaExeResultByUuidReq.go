package cainiaocntec

import (
	"sync"
)

// RpaExeResultByUuidReq 结构体
type RpaExeResultByUuidReq struct {
	// 执行结果code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 详细信息
	Data string `json:"data,omitempty" xml:"data,omitempty"`
	// 访问业务secretKey
	SecretKey string `json:"secret_key,omitempty" xml:"secret_key,omitempty"`
	// 描述
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// uuid
	Uuid string `json:"uuid,omitempty" xml:"uuid,omitempty"`
}

var poolRpaExeResultByUuidReq = sync.Pool{
	New: func() any {
		return new(RpaExeResultByUuidReq)
	},
}

// GetRpaExeResultByUuidReq() 从对象池中获取RpaExeResultByUuidReq
func GetRpaExeResultByUuidReq() *RpaExeResultByUuidReq {
	return poolRpaExeResultByUuidReq.Get().(*RpaExeResultByUuidReq)
}

// ReleaseRpaExeResultByUuidReq 释放RpaExeResultByUuidReq
func ReleaseRpaExeResultByUuidReq(v *RpaExeResultByUuidReq) {
	v.Code = ""
	v.Data = ""
	v.SecretKey = ""
	v.Message = ""
	v.Uuid = ""
	poolRpaExeResultByUuidReq.Put(v)
}
