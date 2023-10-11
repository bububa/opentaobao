package idle

// IdleItemAutoRechargeBatchRemoveResult 结构体
type IdleItemAutoRechargeBatchRemoveResult struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误说明
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 移除失败商品id
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
}
