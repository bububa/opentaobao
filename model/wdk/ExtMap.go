package wdk

// ExtMap 结构体
type ExtMap struct {
	// 订单小号
	OrderNo string `json:"order_no,omitempty" xml:"order_no,omitempty"`
	// 最晚拣货完成时间
	LatestPrepareTime string `json:"latest_prepare_time,omitempty" xml:"latest_prepare_time,omitempty"`
}
