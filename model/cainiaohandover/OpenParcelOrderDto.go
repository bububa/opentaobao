package cainiaohandover

// OpenParcelOrderDto 结构体
type OpenParcelOrderDto struct {
	// 小包状态code
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 小包异常码
	ExceptionCode string `json:"exception_code,omitempty" xml:"exception_code,omitempty"`
	// 小包物流订单编码
	OrderCode string `json:"order_code,omitempty" xml:"order_code,omitempty"`
	// 小包状态名称
	StatusName string `json:"status_name,omitempty" xml:"status_name,omitempty"`
}
