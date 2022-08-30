package ascp

// HiErpDeliverResp 结构体
type HiErpDeliverResp struct {
	// SCP单号
	ScpCode string `json:"scp_code,omitempty" xml:"scp_code,omitempty"`
	// 外部订单号
	OutOrderCode string `json:"out_order_code,omitempty" xml:"out_order_code,omitempty"`
}
