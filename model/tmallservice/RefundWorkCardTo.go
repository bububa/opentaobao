package tmallservice

// RefundWorkCardTo 
type RefundWorkCardTo struct {

    // existingFailure
    
    ExistingFailure   bool `json:"existing_failure,omitempty" xml:"existing_failure,omitempty"`
    

    // failureList
    
    FailureList   []int64 `json:"failure_list,omitempty" xml:"failure_list>int64,omitempty"`
    

    // notExistingList
    
    NotExistingList   []int64 `json:"not_existing_list,omitempty" xml:"not_existing_list>int64,omitempty"`
    

    // refundList
    
    RefundList   []int64 `json:"refund_list,omitempty" xml:"refund_list>int64,omitempty"`
    

}
