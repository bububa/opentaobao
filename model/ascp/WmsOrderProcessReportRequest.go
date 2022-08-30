package ascp

// WmsOrderProcessReportRequest 结构体
type WmsOrderProcessReportRequest struct {
	// (创建发货单)条件必填
	OrderLines []OrderLine `json:"order_lines,omitempty" xml:"order_lines>order_line,omitempty"`
	// 业务请求ID，用于做幂等
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 订单标记 ，用字符串格式来表示订单标记列表
	OrderFlag string `json:"order_flag,omitempty" xml:"order_flag,omitempty"`
	// ERP调翱象接口创建货品的时间戳
	RequestTime int64 `json:"request_time,omitempty" xml:"request_time,omitempty"`
	// 单据信息
	Order *Order `json:"order,omitempty" xml:"order,omitempty"`
	// 单据作业信息
	Process *Process `json:"process,omitempty" xml:"process,omitempty"`
}
