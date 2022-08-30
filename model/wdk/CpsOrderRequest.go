package wdk

// CpsOrderRequest 结构体
type CpsOrderRequest struct {
	// 订单更新开始时间
	StartTime string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 订单更新结束时间
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 页序号，从0开始
	PageIndex int64 `json:"page_index,omitempty" xml:"page_index,omitempty"`
	// 单页大小，不超过200
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
}
