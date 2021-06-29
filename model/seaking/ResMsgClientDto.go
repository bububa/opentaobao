package seaking

// ResMsgClientDTO 
type ResMsgClientDTO struct {
    // 考试结果
    List   []AlibabaAlifanyiMarketExamData `json:"list,omitempty" xml:"list>alibaba_alifanyi_market_exam_data,omitempty"`
    // 错误信息
    ErrorMessage   string `json:"error_message,omitempty" xml:"error_message,omitempty"`
    // 是否成功
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
    // 错误码
    BizErrorCode   int64 `json:"biz_error_code,omitempty" xml:"biz_error_code,omitempty"`
    // 开放ids
    DataList   []string `json:"data_list,omitempty" xml:"data_list>string,omitempty"`
}
