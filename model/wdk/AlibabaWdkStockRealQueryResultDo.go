package wdk

// AlibabaWdkStockRealQueryResultDo 
/* model for simplify = false
type AlibabaWdkStockRealQueryResultDo struct {

    // 错误信息
    
    ErrMsg   string `json:"err_msg,omitempty"`
    

    // 错误码
    
    ErrCode   int64 `json:"err_code,omitempty"`
    

    // 是否成功
    
    Success   bool `json:"success,omitempty"`
    

    // result
    
    Data  *struct {
        InventoryTopResultBo  *InventoryTopResultBo `json:"inventory_top_result_bo,omitempty"`
    } `json:"data,omitempty"`
    

}
*/

// AlibabaWdkStockRealQueryResultDo 
type AlibabaWdkStockRealQueryResultDo struct {

    // 错误信息
    ErrMsg   string `json:"err_msg,omitempty"`

    // 错误码
    ErrCode   int64 `json:"err_code,omitempty"`

    // 是否成功
    Success   bool `json:"success,omitempty"`

    // result
    Data   *InventoryTopResultBo `json:"data,omitempty"`

}
