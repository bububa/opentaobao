package ascp

// DeliveryDecisionQueryRequest 结构体
type DeliveryDecisionQueryRequest struct {
	// 批量（最多可以支持50条）
	DeliveryDecision []DeliveryDecision `json:"delivery_decision,omitempty" xml:"delivery_decision>delivery_decision,omitempty"`
	// 业务请求ID，用于做幂等
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// ERP调翱象接口创建货品的时间戳
	RequestTime int64 `json:"request_time,omitempty" xml:"request_time,omitempty"`
}
