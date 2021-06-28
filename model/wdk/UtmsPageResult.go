package wdk

// UtmsPageResult 
type UtmsPageResult struct {

    // code
    
    Code   string `json:"code,omitempty" xml:"code,omitempty"`
    

    // list
    
    List   []ErpBillDto `json:"list,omitempty" xml:"list,omitempty"`
    

    // msg
    
    Msg   string `json:"msg,omitempty" xml:"msg,omitempty"`
    

    // success
    
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    

    // totalCount
    
    TotalCount   int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
    

}
