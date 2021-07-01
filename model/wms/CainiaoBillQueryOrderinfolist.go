package wms

// CainiaoBillQueryOrderinfolist 结构体
type CainiaoBillQueryOrderinfolist struct {
	// 订单信息
	OrderInfo *CainiaoBillQueryOrderinfo `json:"order_info,omitempty" xml:"order_info,omitempty"`
}
