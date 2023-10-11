package drugtrace

// CodeScanDto 结构体
type CodeScanDto struct {
	// 扫码时间
	QueryDate string `json:"query_date,omitempty" xml:"query_date,omitempty"`
	// 扫码来源：支付宝，淘宝，天猫，未知
	QueryAppNameFormat string `json:"query_app_name_format,omitempty" xml:"query_app_name_format,omitempty"`
}
