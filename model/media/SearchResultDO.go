package media

// SearchResultDO 
/* model for simplify = false
type SearchResultDO struct {

    // 错误码
    
    ResultCode   string `json:"result_code,omitempty"`
    

    // 视频信息列表
    
    ResultList  struct {
        VideoItemExtDO  []VideoItemExtDO `json:"video_item_ext_do,omitempty"`
    } `json:"result_list,omitempty"`
    

    // 调用是否成功
    
    IsSuccess   bool `json:"is_success,omitempty"`
    

    // 总视频数
    
    TotalNum   int64 `json:"total_num,omitempty"`
    

}
*/

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
