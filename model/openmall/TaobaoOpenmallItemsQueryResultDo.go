package openmall

// TaobaoOpenmallItemsQueryResultDO 
type TaobaoOpenmallItemsQueryResultDO struct {
    // 错误码
    ErrorCode   string `json:"error_code,omitempty" xml:"error_code,omitempty"`
    // 错误信息
    ErrorMsg   string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
    // 商品列表
    ItemList   []TopItemVo `json:"item_list,omitempty" xml:"item_list>top_item_vo,omitempty"`
    // 调用是否成功
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
}
