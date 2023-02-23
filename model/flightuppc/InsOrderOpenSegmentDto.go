package flightuppc

// InsOrderOpenSegmentDto 结构体
type InsOrderOpenSegmentDto struct {
	// 航司名称
	CompanyName string `json:"company_name,omitempty" xml:"company_name,omitempty"`
	// 到达城市
	ArrCity string `json:"arr_city,omitempty" xml:"arr_city,omitempty"`
	// 出发城市
	DepCity string `json:"dep_city,omitempty" xml:"dep_city,omitempty"`
	// 子保单号
	PolicyNo string `json:"policy_no,omitempty" xml:"policy_no,omitempty"`
	// 航段编号
	SegmentNo string `json:"segment_no,omitempty" xml:"segment_no,omitempty"`
	// 出发时间
	StartTime string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 到达时间
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// attribute
	Attribute string `json:"attribute,omitempty" xml:"attribute,omitempty"`
	// 保险订单号
	TcOrderId int64 `json:"tc_order_id,omitempty" xml:"tc_order_id,omitempty"`
	// 外部订单号
	OutOrderId int64 `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
}
