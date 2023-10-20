package ascp

import (
	"sync"
)

// HiErpCloseResp 结构体
type HiErpCloseResp struct {
	// 外部订单号
	OutBizId string `json:"out_biz_id,omitempty" xml:"out_biz_id,omitempty"`
	// 履约单号
	OrderCode string `json:"order_code,omitempty" xml:"order_code,omitempty"`
	// 取消任务创建结果
	TaskExcMsg string `json:"task_exc_msg,omitempty" xml:"task_exc_msg,omitempty"`
}

var poolHiErpCloseResp = sync.Pool{
	New: func() any {
		return new(HiErpCloseResp)
	},
}

// GetHiErpCloseResp() 从对象池中获取HiErpCloseResp
func GetHiErpCloseResp() *HiErpCloseResp {
	return poolHiErpCloseResp.Get().(*HiErpCloseResp)
}

// ReleaseHiErpCloseResp 释放HiErpCloseResp
func ReleaseHiErpCloseResp(v *HiErpCloseResp) {
	v.OutBizId = ""
	v.OrderCode = ""
	v.TaskExcMsg = ""
	poolHiErpCloseResp.Put(v)
}
