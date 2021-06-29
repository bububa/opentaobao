package mozi

// PageTenantSubAdminsResult 
type PageTenantSubAdminsResult struct {

    // 总条数
    
    TotalSize   int64 `json:"total_size,omitempty" xml:"total_size,omitempty"`
    

    // data数据
    
    Datas   []TenantAdmin `json:"datas,omitempty" xml:"datas,omitempty"`
    

    // success
    
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    

    // requestId
    
    RequestId   string `json:"request_id,omitempty" xml:"request_id,omitempty"`
    

    // 页的条目数
    
    PageSize   int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
    

    // 当前页数
    
    CurrentPage   int64 `json:"current_page,omitempty" xml:"current_page,omitempty"`
    

    // responseMessage
    
    ResponseMessage   string `json:"response_message,omitempty" xml:"response_message,omitempty"`
    

    // responseCode
    
    ResponseCode   string `json:"response_code,omitempty" xml:"response_code,omitempty"`
    

}
