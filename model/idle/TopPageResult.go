package idle

// TopPageResult 
type TopPageResult struct {
    // 商品列表
    ItemList   []IdleItemApiDO `json:"item_list,omitempty" xml:"item_list>idle_item_api_do,omitempty"`
    // 错误码
    ErrCode   string `json:"err_code,omitempty" xml:"err_code,omitempty"`
    // 错误描述
    ErrMsg   string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
    // 是否有下一页
    NextPage   bool `json:"next_page,omitempty" xml:"next_page,omitempty"`
    // 是否成功
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    // 总结果数
    Total   int64 `json:"total,omitempty" xml:"total,omitempty"`
}
