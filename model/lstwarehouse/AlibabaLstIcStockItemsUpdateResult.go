package lstwarehouse

// AlibabaLstIcStockItemsUpdateResult 
type AlibabaLstIcStockItemsUpdateResult struct {

    // 是否成功
    
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    

    // errorMsg
    
    ErrorMessage   string `json:"error_message,omitempty" xml:"error_message,omitempty"`
    

    // 错误码
    
    ErrorCode   string `json:"error_code,omitempty" xml:"error_code,omitempty"`
    

    // 列表
    
    ContentList   []Content `json:"content_list,omitempty" xml:"content_list,omitempty"`
    

}
