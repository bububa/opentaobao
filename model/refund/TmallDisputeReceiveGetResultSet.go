package refund

// TmallDisputeReceiveGetResultSet 
type TmallDisputeReceiveGetResultSet struct {

    // 当前页面的纠纷单数量
    PageResults   int64 `json:"page_results,omitempty"`

    // 所有符合查询条件的纠纷单数量
    TotalResults   int64 `json:"total_results,omitempty"`

    // 是否还有下一页
    HasNext   bool `json:"has_next,omitempty"`

    // 所抛出异常
    Exception   string `json:"exception,omitempty"`

    // 错误码
    ErrorCode   string `json:"error_code,omitempty"`

    // 错误信息
    ErrorMsg   string `json:"error_msg,omitempty"`

    // results
    Results   []Dispute `json:"results,omitempty"`

}
