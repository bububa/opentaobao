package tmallservice

// RefundWorkCardTo 
/* model for simplify = false
type RefundWorkCardTo struct {

    // existingFailure
    
    ExistingFailure   bool `json:"existing_failure,omitempty"`
    

    // failureList
    
    FailureList  struct {
        Number  []int64 `json:"int64,omitempty"`
    } `json:"failure_list,omitempty"`
    

    // notExistingList
    
    NotExistingList  struct {
        Number  []int64 `json:"int64,omitempty"`
    } `json:"not_existing_list,omitempty"`
    

    // refundList
    
    RefundList  struct {
        Number  []int64 `json:"int64,omitempty"`
    } `json:"refund_list,omitempty"`
    

}
*/

// RefundWorkCardTo 
type RefundWorkCardTo struct {

    // existingFailure
    ExistingFailure   bool `json:"existing_failure,omitempty"`

    // failureList
    FailureList   []int64 `json:"failure_list,omitempty"`

    // notExistingList
    NotExistingList   []int64 `json:"not_existing_list,omitempty"`

    // refundList
    RefundList   []int64 `json:"refund_list,omitempty"`

}
