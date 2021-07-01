package waybill

// WaybillApplyPrintCheckInfo 结构体
type WaybillApplyPrintCheckInfo struct {
	// 打印提示信息编码
	NoticeCode string `json:"notice_code,omitempty" xml:"notice_code,omitempty"`
	// 打印次数
	PrintQuantity int64 `json:"print_quantity,omitempty" xml:"print_quantity,omitempty"`
	// 打印提示信息
	NoticeMessage string `json:"notice_message,omitempty" xml:"notice_message,omitempty"`
	// 电子面单号
	WaybillCode string `json:"waybill_code,omitempty" xml:"waybill_code,omitempty"`
}
