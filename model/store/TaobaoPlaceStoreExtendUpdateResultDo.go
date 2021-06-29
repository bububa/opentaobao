package store

// TaobaoPlaceStoreExtendUpdateResultDo 
type TaobaoPlaceStoreExtendUpdateResultDo struct {

    // 错误信息
    
    ErrorMsg   string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
    

    // 是否失败
    
    Failured   bool `json:"failured,omitempty" xml:"failured,omitempty"`
    

    // 完整错误信息
    
    FullErrorMsg   string `json:"full_error_msg,omitempty" xml:"full_error_msg,omitempty"`
    

    // 模型
    
    Models   *Models `json:"models,omitempty" xml:"models,omitempty"`
    

    // 关键主键
    
    PriKey   string `json:"pri_key,omitempty" xml:"pri_key,omitempty"`
    

    // 结果
    
    Result   bool `json:"result,omitempty" xml:"result,omitempty"`
    

    // 错误码
    
    ResultCode   string `json:"result_code,omitempty" xml:"result_code,omitempty"`
    

    // 调用是否成功
    
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    

    // 总数量
    
    TotalNum   int64 `json:"total_num,omitempty" xml:"total_num,omitempty"`
    

}
