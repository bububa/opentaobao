package wdk

// ApiPageResult 
type ApiPageResult struct {

    // 是否还有下一页
    
    HasNext   bool `json:"has_next,omitempty" xml:"has_next,omitempty"`
    

    // 页码
    
    PageIndex   int64 `json:"page_index,omitempty" xml:"page_index,omitempty"`
    

    // 每页记录数
    
    PageSize   int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
    

    // 错误编码
    
    ErrCode   string `json:"err_code,omitempty" xml:"err_code,omitempty"`
    

    // 返回内容
    
    Models   []CouponStatisticsResultDO `json:"models,omitempty" xml:"models,omitempty"`
    

    // 错误信息
    
    ErrMsg   string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
    

    // 是否成功
    
    Success   string `json:"success,omitempty" xml:"success,omitempty"`
    

}
