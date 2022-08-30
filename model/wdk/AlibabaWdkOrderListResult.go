package wdk

// AlibabaWdkOrderListResult 结构体
type AlibabaWdkOrderListResult struct {
	// 订单列表
	Orders []Order `json:"orders,omitempty" xml:"orders>order,omitempty"`
	// 接口返回码。如果返回 HM02008888888001 代表成功，其他值代表失败。调用方需要根据返回码判断，失败重试
	ReturnCode string `json:"return_code,omitempty" xml:"return_code,omitempty"`
	// 接口返回码描述
	ReturnMsg string `json:"return_msg,omitempty" xml:"return_msg,omitempty"`
	// 返回下一查询页的序号。如果返回值是-1，则无下一页。数据拉取完成。
	NextIndex int64 `json:"next_index,omitempty" xml:"next_index,omitempty"`
	// 返回本查询条件下的数据总数。仅在传入page_index=0时返回,在其他情况下返回0
	TotalNumber int64 `json:"total_number,omitempty" xml:"total_number,omitempty"`
}
