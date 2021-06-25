package promotion

// CouponApplyResult 
type CouponApplyResult struct {

    // 请求唯一id，问题排查
    TraceId   string `json:"trace_id,omitempty"`

    // 领取结果，领取成功为true，否则为false
    ApplySuccess   bool `json:"apply_success,omitempty"`

}
