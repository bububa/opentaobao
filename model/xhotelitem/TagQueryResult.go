package xhotelitem

// TagQueryResult 
type TagQueryResult struct {
    // 总数
    TotalAmount   int64 `json:"total_amount,omitempty" xml:"total_amount,omitempty"`
    // 是否成功
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    // 列表
    TagEntityDoList   []TagEntityDoList `json:"tag_entity_do_list,omitempty" xml:"tag_entity_do_list>tag_entity_do_list,omitempty"`
    // token
    TokenStr   string `json:"token_str,omitempty" xml:"token_str,omitempty"`
    // 耗时
    SpentTime   int64 `json:"spent_time,omitempty" xml:"spent_time,omitempty"`
    // 错误信息
    ErrorMsg   string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
}
