package ieagency

// QueryChangeAgentListRs 
type QueryChangeAgentListRs struct {
    // 错误码
    ApiErrorCode   int64 `json:"api_error_code,omitempty" xml:"api_error_code,omitempty"`
    // 错误信息
    ApiErrorMsg   string `json:"api_error_msg,omitempty" xml:"api_error_msg,omitempty"`
    // 改签单信息
    ChangeOrderDOs   []IeChangeOrderVo `json:"change_order_d_os,omitempty" xml:"change_order_d_os>ie_change_order_vo,omitempty"`
    // 分页信息
    PageInfoDO   *BasePageDO `json:"page_info_d_o,omitempty" xml:"page_info_d_o,omitempty"`
    // 是否成功
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
}
