package tmallservice

// WorkcardBaseRequest 
/* model for simplify = false
type WorkcardBaseRequest struct {

    // 请求来源账号
    
    RequestSource  *struct {
        WorkcardRequestSource  *WorkcardRequestSource `json:"workcard_request_source,omitempty"`
    } `json:"request_source,omitempty"`
    

    // 工单ID
    
    WorkcardId   int64 `json:"workcard_id,omitempty"`
    

}
*/

// WorkcardBaseRequest 
type WorkcardBaseRequest struct {

    // 请求来源账号
    RequestSource   *WorkcardRequestSource `json:"request_source,omitempty"`

    // 工单ID
    WorkcardId   int64 `json:"workcard_id,omitempty"`

}
