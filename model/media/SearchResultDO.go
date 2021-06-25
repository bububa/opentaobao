package media

// SearchResultDO 
type SearchResultDO struct {

    // 错误码
    ResultCode   string `json:"result_code,omitempty"`

    // 视频信息列表
    ResultList   []VideoItemExtDO `json:"result_list,omitempty"`

    // 调用是否成功
    IsSuccess   bool `json:"is_success,omitempty"`

    // 总视频数
    TotalNum   int64 `json:"total_num,omitempty"`

}
