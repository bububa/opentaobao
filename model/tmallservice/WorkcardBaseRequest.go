package tmallservice

// WorkcardBaseRequest 结构体
type WorkcardBaseRequest struct {
	// 请求来源账号
	RequestSource *WorkcardRequestSource `json:"request_source,omitempty" xml:"request_source,omitempty"`
	// 工单ID
	WorkcardId int64 `json:"workcard_id,omitempty" xml:"workcard_id,omitempty"`
}
