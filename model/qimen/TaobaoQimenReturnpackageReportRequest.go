package qimen

// TaobaoQimenReturnpackageReportRequest 结构体
type TaobaoQimenReturnpackageReportRequest struct {
	// 订单
	Order *Order `json:"order,omitempty" xml:"order,omitempty"`
	// 包裹列表
	Packages *Packages `json:"packages,omitempty" xml:"packages,omitempty"`
}
