package tmallsc

// BuyerRefuseAcceptRequest 
type BuyerRefuseAcceptRequest struct {
    // 工单ID
    WorkcardId   int64 `json:"workcard_id,omitempty" xml:"workcard_id,omitempty"`
    // 拒收备注
    RefuseMemo   string `json:"refuse_memo,omitempty" xml:"refuse_memo,omitempty"`
}
