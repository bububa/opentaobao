package wdk

// OrderListSyncPagedResult 结构体
type OrderListSyncPagedResult struct {
	// orders
	Orders []OrderSyncDto `json:"orders,omitempty" xml:"orders>order_sync_dto,omitempty"`
	// returnMsg
	ReturnMsg string `json:"return_msg,omitempty" xml:"return_msg,omitempty"`
	// returnCode
	ReturnCode string `json:"return_code,omitempty" xml:"return_code,omitempty"`
	// 返回订单总数量
	TotalNumber int64 `json:"total_number,omitempty" xml:"total_number,omitempty"`
	// 返回下一查询页的序号。如果返回值是-1，则无下一页。数据拉取完成。
	NextIndex int64 `json:"next_index,omitempty" xml:"next_index,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
