package tbtrade

// AppBillQueryRequest 结构体
type AppBillQueryRequest struct {
	// 账单时间
	BizDate string `json:"biz_date,omitempty" xml:"biz_date,omitempty"`
	// 页号，从1开始
	PageNo int64 `json:"page_no,omitempty" xml:"page_no,omitempty"`
	// 页大小，不得超过1000
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
}
