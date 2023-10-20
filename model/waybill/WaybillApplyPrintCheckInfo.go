package waybill

import (
	"sync"
)

// WaybillApplyPrintCheckInfo 结构体
type WaybillApplyPrintCheckInfo struct {
	// 打印提示信息编码
	NoticeCode string `json:"notice_code,omitempty" xml:"notice_code,omitempty"`
	// 打印提示信息
	NoticeMessage string `json:"notice_message,omitempty" xml:"notice_message,omitempty"`
	// 电子面单号
	WaybillCode string `json:"waybill_code,omitempty" xml:"waybill_code,omitempty"`
	// 打印次数
	PrintQuantity int64 `json:"print_quantity,omitempty" xml:"print_quantity,omitempty"`
}

var poolWaybillApplyPrintCheckInfo = sync.Pool{
	New: func() any {
		return new(WaybillApplyPrintCheckInfo)
	},
}

// GetWaybillApplyPrintCheckInfo() 从对象池中获取WaybillApplyPrintCheckInfo
func GetWaybillApplyPrintCheckInfo() *WaybillApplyPrintCheckInfo {
	return poolWaybillApplyPrintCheckInfo.Get().(*WaybillApplyPrintCheckInfo)
}

// ReleaseWaybillApplyPrintCheckInfo 释放WaybillApplyPrintCheckInfo
func ReleaseWaybillApplyPrintCheckInfo(v *WaybillApplyPrintCheckInfo) {
	v.NoticeCode = ""
	v.NoticeMessage = ""
	v.WaybillCode = ""
	v.PrintQuantity = 0
	poolWaybillApplyPrintCheckInfo.Put(v)
}
