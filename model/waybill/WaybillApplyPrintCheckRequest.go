package waybill

import (
	"sync"
)

// WaybillApplyPrintCheckRequest 结构体
type WaybillApplyPrintCheckRequest struct {
	// 面单详情信息
	PrintCheckInfoCols []PrintCheckInfo `json:"print_check_info_cols,omitempty" xml:"print_check_info_cols>print_check_info,omitempty"`
	// 物流服务商Code
	CpCode string `json:"cp_code,omitempty" xml:"cp_code,omitempty"`
}

var poolWaybillApplyPrintCheckRequest = sync.Pool{
	New: func() any {
		return new(WaybillApplyPrintCheckRequest)
	},
}

// GetWaybillApplyPrintCheckRequest() 从对象池中获取WaybillApplyPrintCheckRequest
func GetWaybillApplyPrintCheckRequest() *WaybillApplyPrintCheckRequest {
	return poolWaybillApplyPrintCheckRequest.Get().(*WaybillApplyPrintCheckRequest)
}

// ReleaseWaybillApplyPrintCheckRequest 释放WaybillApplyPrintCheckRequest
func ReleaseWaybillApplyPrintCheckRequest(v *WaybillApplyPrintCheckRequest) {
	v.PrintCheckInfoCols = v.PrintCheckInfoCols[:0]
	v.CpCode = ""
	poolWaybillApplyPrintCheckRequest.Put(v)
}
