package tmallcar

// TopOrderQuery 结构体
type TopOrderQuery struct {
	// 开始时间
	StartTime string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 结束时间
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 当前页
	PageNo int64 `json:"page_no,omitempty" xml:"page_no,omitempty"`
	// 每页个数
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 订单号
	OrderId int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
}
