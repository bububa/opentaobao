package paimai

// VehicleDetectServerReport4Top 结构体
type VehicleDetectServerReport4Top struct {
	// 拍卖服务单单号（与检测单单号不能同时为空）
	ServiceCaseNo string `json:"service_case_no,omitempty" xml:"service_case_no,omitempty"`
	// 检测机构检测单单号（与拍卖服务单单号，不能同时为空）
	OrderNum string `json:"order_num,omitempty" xml:"order_num,omitempty"`
}
