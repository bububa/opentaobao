package tmallservice

// AcceptWorkcardRequest 结构体
type AcceptWorkcardRequest struct {
	// 工单号
	WorkcardId int64 `json:"workcard_id,omitempty" xml:"workcard_id,omitempty"`
}
