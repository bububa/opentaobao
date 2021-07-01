package cainiaoecc

// DelayExceptionCountDto 
type DelayExceptionCountDto struct {
    // 异常总数
    ExceptionNum   int64 `json:"exception_num,omitempty" xml:"exception_num,omitempty"`
    // 商家Id
    SellerId   int64 `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
    // CP异常回复总数
    CpReplyExceptionNum   int64 `json:"cp_reply_exception_num,omitempty" xml:"cp_reply_exception_num,omitempty"`
    // 扩展字段
    ExtendFields   string `json:"extend_fields,omitempty" xml:"extend_fields,omitempty"`
}
