package ascpchannel

// WaybillCloudPrintDto 结构体
type WaybillCloudPrintDto struct {
	// 对应传入的packageCode
	PackageCode string `json:"package_code,omitempty" xml:"package_code,omitempty"`
	// 面单号
	WaybillCode string `json:"waybill_code,omitempty" xml:"waybill_code,omitempty"`
	// 父面单号 本次没有，为子母单预留
	ParentWaybillCode string `json:"parent_waybill_code,omitempty" xml:"parent_waybill_code,omitempty"`
	// 面单信息
	PrintData string `json:"print_data,omitempty" xml:"print_data,omitempty"`
}
