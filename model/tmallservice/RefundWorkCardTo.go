package tmallservice

// RefundWorkCardTo 
type RefundWorkCardTo struct {

    // existingFailure
    ExistingFailure   bool `json:"existing_failure,omitempty"`

    // failureList
    FailureList   []Number `json:"failure_list,omitempty"`

    // notExistingList
    NotExistingList   []Number `json:"not_existing_list,omitempty"`

    // refundList
    RefundList   []Number `json:"refund_list,omitempty"`

}
