package ieagency

// IeRefundReasonDo 
type IeRefundReasonDo struct {
    // 原因描述
    Reason   string `json:"reason,omitempty" xml:"reason,omitempty"`
    // 原因类型
    ReasonType   int64 `json:"reason_type,omitempty" xml:"reason_type,omitempty"`
}
