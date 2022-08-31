package aesolution

// RmaStateUpdateRequest 结构体
type RmaStateUpdateRequest struct {
	// RMA&#39;s ID
	RmaId string `json:"rma_id,omitempty" xml:"rma_id,omitempty"`
	// RMA&#39;s state information
	RmaState *RmaState `json:"rma_state,omitempty" xml:"rma_state,omitempty"`
}
