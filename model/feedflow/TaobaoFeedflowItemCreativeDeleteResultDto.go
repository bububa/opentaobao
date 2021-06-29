package feedflow

// TaobaoFeedflowItemCreativeDeleteResultDto 
type TaobaoFeedflowItemCreativeDeleteResultDto struct {
    // 消息
    Message   string `json:"message,omitempty" xml:"message,omitempty"`
    // 调用是否成功,true-成功，false-失败
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    // 错误信息
    ErrorList   []ErrorInfoDto `json:"error_list,omitempty" xml:"error_list>error_info_dto,omitempty"`
}
