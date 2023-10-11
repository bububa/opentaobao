package idle

// TopListResult 结构体
type TopListResult struct {
	// 移除失败商品信息列表
	Data []IdleItemAutoRechargeBatchRemoveResult `json:"data,omitempty" xml:"data>idle_item_auto_recharge_batch_remove_result,omitempty"`
	// 错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 错误说明
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 请求是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
