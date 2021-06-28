package wdk

// PromotionPriceResult 
type PromotionPriceResult struct {

    // 页码
    
    PageIndex   int64 `json:"page_index,omitempty" xml:"page_index,omitempty"`
    

    // 单页数据大小
    
    PageSize   int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
    

    // 总页数
    
    PageCount   int64 `json:"page_count,omitempty" xml:"page_count,omitempty"`
    

    // 总数量
    
    Total   int64 `json:"total,omitempty" xml:"total,omitempty"`
    

    // 促销信息记录
    
    ItemList   []PromotionPriceDO `json:"item_list,omitempty" xml:"item_list,omitempty"`
    

    // 错误编码
    
    ErrCode   string `json:"err_code,omitempty" xml:"err_code,omitempty"`
    

    // 错误信息
    
    ErrMsg   string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
    

    // 是否成功
    
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    

}
