package tbk

// TopApiRefundRptOption 结构体
type TopApiRefundRptOption struct {
	// 开始时间
	Starttime string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// pagesize
	Pagesize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 1-维权发起时间，2-订单结算时间（正向订单），3-维权完成时间，4-订单创建时间，5-订单更新时间
	Searchtype int64 `json:"search_type,omitempty" xml:"search_type,omitempty"`
	// 1 表示2方，2表示3方，0表示不限
	Refundtype int64 `json:"refund_type,omitempty" xml:"refund_type,omitempty"`
	// pagenumber
	Pageno int64 `json:"page_no,omitempty" xml:"page_no,omitempty"`
	// 1代表渠道关系id，2代表会员关系id
	Biztype int64 `json:"biz_type,omitempty" xml:"biz_type,omitempty"`
}
