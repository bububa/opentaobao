package wdk

// UtmsPageResult 
/* model for simplify = false
type UtmsPageResult struct {

    // code
    
    Code   string `json:"code,omitempty"`
    

    // list
    
    List  struct {
        ErpBillDto  []ErpBillDto `json:"erp_bill_dto,omitempty"`
    } `json:"list,omitempty"`
    

    // msg
    
    Msg   string `json:"msg,omitempty"`
    

    // success
    
    Success   bool `json:"success,omitempty"`
    

    // totalCount
    
    TotalCount   int64 `json:"total_count,omitempty"`
    

}
*/

// UtmsPageResult 
type UtmsPageResult struct {

    // code
    Code   string `json:"code,omitempty"`

    // list
    List   []ErpBillDto `json:"list,omitempty"`

    // msg
    Msg   string `json:"msg,omitempty"`

    // success
    Success   bool `json:"success,omitempty"`

    // totalCount
    TotalCount   int64 `json:"total_count,omitempty"`

}
