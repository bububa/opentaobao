package wdk

// AlibabaWdkStockCabinetQueryResultDo 
type AlibabaWdkStockCabinetQueryResultDo struct {

    // 错误信息详细描述
    
    ErrMsg   string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
    

    // 错误code
    
    ErrCode   int64 `json:"err_code,omitempty" xml:"err_code,omitempty"`
    

    // 调用是否成功
    
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    

    // 返回的结果数据
    
    Data   *InventoryTopResultBo `json:"data,omitempty" xml:"data,omitempty"`
    

}
