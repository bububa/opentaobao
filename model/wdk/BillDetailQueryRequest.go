package wdk

// BillDetailQueryRequest 结构体
type BillDetailQueryRequest struct {
	// 经营店ID
	StoreId string `json:"store_id,omitempty" xml:"store_id,omitempty"`
	// 查询的页码，从1开始
	Current int64 `json:"current,omitempty" xml:"current,omitempty"`
	// 页大小
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 账单起始时间: yyyyMMdd
	BillStartDate int64 `json:"bill_start_date,omitempty" xml:"bill_start_date,omitempty"`
	// 账单结束时间: yyyyMMdd
	BillEndDate int64 `json:"bill_end_date,omitempty" xml:"bill_end_date,omitempty"`
}
