package waybill

// WaybillApplyPrintCheckRequest 结构体
type WaybillApplyPrintCheckRequest struct {
	// 面单详情信息
	PrintCheckInfoCols []PrintCheckInfo `json:"print_check_info_cols,omitempty" xml:"print_check_info_cols>print_check_info,omitempty"`
	// 物流服务商Code
	CpCode string `json:"cp_code,omitempty" xml:"cp_code,omitempty"`
}
