package tmallservice

// UpdateAttributeRequest 
type UpdateAttributeRequest struct {

    // 服务回访记录信息
    
    CallRecord   string `json:"call_record,omitempty" xml:"call_record,omitempty"`
    

    // 请求来源
    
    RequestSource   *WorkcardRequestSource `json:"request_source,omitempty" xml:"request_source,omitempty"`
    

    // 工单ID
    
    WorkcardId   int64 `json:"workcard_id,omitempty" xml:"workcard_id,omitempty"`
    

}
