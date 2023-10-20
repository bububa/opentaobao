package tmallservice

import (
	"sync"
)

// TmallServicecenterServicestoreCreateResult 结构体
type TmallServicecenterServicestoreCreateResult struct {
	// 错误码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 错误信息
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 门店数据
	ResultData *ServiceStoreCreateResp `json:"result_data,omitempty" xml:"result_data,omitempty"`
	// 调用是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTmallServicecenterServicestoreCreateResult = sync.Pool{
	New: func() any {
		return new(TmallServicecenterServicestoreCreateResult)
	},
}

// GetTmallServicecenterServicestoreCreateResult() 从对象池中获取TmallServicecenterServicestoreCreateResult
func GetTmallServicecenterServicestoreCreateResult() *TmallServicecenterServicestoreCreateResult {
	return poolTmallServicecenterServicestoreCreateResult.Get().(*TmallServicecenterServicestoreCreateResult)
}

// ReleaseTmallServicecenterServicestoreCreateResult 释放TmallServicecenterServicestoreCreateResult
func ReleaseTmallServicecenterServicestoreCreateResult(v *TmallServicecenterServicestoreCreateResult) {
	v.MsgCode = ""
	v.MsgInfo = ""
	v.ResultData = nil
	v.Success = false
	poolTmallServicecenterServicestoreCreateResult.Put(v)
}
