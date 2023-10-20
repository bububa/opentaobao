package wlb

import (
	"sync"
)

// WlbWmsInventoryLackUploadResp 结构体
type WlbWmsInventoryLackUploadResp struct {
	// 错误编码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolWlbWmsInventoryLackUploadResp = sync.Pool{
	New: func() any {
		return new(WlbWmsInventoryLackUploadResp)
	},
}

// GetWlbWmsInventoryLackUploadResp() 从对象池中获取WlbWmsInventoryLackUploadResp
func GetWlbWmsInventoryLackUploadResp() *WlbWmsInventoryLackUploadResp {
	return poolWlbWmsInventoryLackUploadResp.Get().(*WlbWmsInventoryLackUploadResp)
}

// ReleaseWlbWmsInventoryLackUploadResp 释放WlbWmsInventoryLackUploadResp
func ReleaseWlbWmsInventoryLackUploadResp(v *WlbWmsInventoryLackUploadResp) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Success = false
	poolWlbWmsInventoryLackUploadResp.Put(v)
}
