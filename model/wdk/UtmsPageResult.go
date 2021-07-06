package wdk

// UtmsPageResult 结构体
type UtmsPageResult struct {
	// list
	List []ErpBillDto `json:"list,omitempty" xml:"list>erp_bill_dto,omitempty"`
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// msg
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// totalCount
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
