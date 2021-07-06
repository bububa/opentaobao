package aesolution

// LogisticOrderStateUpdateRequest 结构体
type LogisticOrderStateUpdateRequest struct {
	// Tracking code
	TrackingCode string `json:"tracking_code,omitempty" xml:"tracking_code,omitempty"`
	// Logistic order state information
	RmaLogisticOrderState *RmaLogisticOrderState `json:"rma_logistic_order_state,omitempty" xml:"rma_logistic_order_state,omitempty"`
}
