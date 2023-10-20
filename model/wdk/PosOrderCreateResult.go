package wdk

import (
	"sync"
)

// PosOrderCreateResult 结构体
type PosOrderCreateResult struct {
	// subOrderDOList
	SubOrderDOList []PosSubOrderResult `json:"sub_order_d_o_list,omitempty" xml:"sub_order_d_o_list>pos_sub_order_result,omitempty"`
	// 结果msg
	ReturnMsg string `json:"return_msg,omitempty" xml:"return_msg,omitempty"`
	// 结果码
	ReturnCode string `json:"return_code,omitempty" xml:"return_code,omitempty"`
	// outOrderId
	OutOrderId string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
	// mainOrderId
	MainOrderId int64 `json:"main_order_id,omitempty" xml:"main_order_id,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolPosOrderCreateResult = sync.Pool{
	New: func() any {
		return new(PosOrderCreateResult)
	},
}

// GetPosOrderCreateResult() 从对象池中获取PosOrderCreateResult
func GetPosOrderCreateResult() *PosOrderCreateResult {
	return poolPosOrderCreateResult.Get().(*PosOrderCreateResult)
}

// ReleasePosOrderCreateResult 释放PosOrderCreateResult
func ReleasePosOrderCreateResult(v *PosOrderCreateResult) {
	v.SubOrderDOList = v.SubOrderDOList[:0]
	v.ReturnMsg = ""
	v.ReturnCode = ""
	v.OutOrderId = ""
	v.MainOrderId = 0
	v.Success = false
	poolPosOrderCreateResult.Put(v)
}
