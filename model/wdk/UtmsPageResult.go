package wdk

// UtmsPageResult 
type UtmsPageResult struct {
    // code
    Code   string `json:"code,omitempty" xml:"code,omitempty"`
    // list
    List   []ErpBillDTO `json:"list,omitempty" xml:"list>erp_bill_dto,omitempty"`
    // msg
    Msg   string `json:"msg,omitempty" xml:"msg,omitempty"`
    // success
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    // totalCount
    TotalCount   int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
}
