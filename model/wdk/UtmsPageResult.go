package wdk

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
