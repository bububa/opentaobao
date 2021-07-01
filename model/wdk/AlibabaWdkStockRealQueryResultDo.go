package wdk

// AlibabaWdkStockRealQueryResultDo 
type AlibabaWdkStockRealQueryResultDo struct {
    // 错误信息
    ErrMsg   string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
    // 错误码
    ErrCode   int64 `json:"err_code,omitempty" xml:"err_code,omitempty"`
    // 是否成功
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    // result
    Data   *InventoryTopResultBo `json:"data,omitempty" xml:"data,omitempty"`
}
