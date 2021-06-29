package aesolution

// LogisticOrderStateUpdateRequest 
type LogisticOrderStateUpdateRequest struct {
    // Logistic order state information
    RmaLogisticOrderState   *RmaLogisticOrderState `json:"rma_logistic_order_state,omitempty" xml:"rma_logistic_order_state,omitempty"`
    // Tracking code
    TrackingCode   string `json:"tracking_code,omitempty" xml:"tracking_code,omitempty"`
}
