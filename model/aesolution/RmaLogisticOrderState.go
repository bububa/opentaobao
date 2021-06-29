package aesolution

// RmaLogisticOrderState 
type RmaLogisticOrderState struct {

    // State date. PST time
    
    StateDate   string `json:"state_date,omitempty" xml:"state_date,omitempty"`
    

    // values CANCELLED, PRODUCT_CAPTURED, INCIDENT, PRODUCT_DELIVERED
    
    State   string `json:"state,omitempty" xml:"state,omitempty"`
    

    // Logistic order detail
    
    OrderStateDetail   string `json:"order_state_detail,omitempty" xml:"order_state_detail,omitempty"`
    

}
