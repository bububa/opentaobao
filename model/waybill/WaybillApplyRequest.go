package waybill

// WaybillApplyRequest 结构体
type WaybillApplyRequest struct {
	// 物流服务商ID
	CpCode string `json:"cp_code,omitempty" xml:"cp_code,omitempty"`
}
