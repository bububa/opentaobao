package wlb

// CloudPrinterPrintRequest 结构体
type CloudPrinterPrintRequest struct {
	// 共享码
	ShareCode string `json:"share_code,omitempty" xml:"share_code,omitempty"`
	// 打印机 id
	Uid string `json:"uid,omitempty" xml:"uid,omitempty"`
	// 自定义内容
	CustomData *CustomData `json:"custom_data,omitempty" xml:"custom_data,omitempty"`
	// 打印数据
	PrintData *PrintData `json:"print_data,omitempty" xml:"print_data,omitempty"`
}
