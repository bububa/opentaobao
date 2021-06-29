package aesolution

// RmaStateUpdateRequest 
type RmaStateUpdateRequest struct {
    // RMA's state information
    RmaState   *RmaState `json:"rma_state,omitempty" xml:"rma_state,omitempty"`
    // RMA's ID
    RmaId   string `json:"rma_id,omitempty" xml:"rma_id,omitempty"`
}
