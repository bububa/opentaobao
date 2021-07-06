package aesolution

// RmaStateUpdateRequest 结构体
type RmaStateUpdateRequest struct {
	// RMA's ID
	RmaId string `json:"rma_id,omitempty" xml:"rma_id,omitempty"`
	// RMA's state information
	RmaState *RmaState `json:"rma_state,omitempty" xml:"rma_state,omitempty"`
}
