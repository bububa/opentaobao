package wdk

// OrderSyncRefundListResult 
type OrderSyncRefundListResult struct {
    // 下一页
    NextIndex   int64 `json:"next_index,omitempty" xml:"next_index,omitempty"`
    // 退款单具体详情
    Orders   []OrderSyncRefundDTO `json:"orders,omitempty" xml:"orders>order_sync_refund_dto,omitempty"`
    // 结果code
    ReturnCode   string `json:"return_code,omitempty" xml:"return_code,omitempty"`
    // 返回的信息
    ReturnMsg   string `json:"return_msg,omitempty" xml:"return_msg,omitempty"`
    // 是否成功
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    // 总数量
    TotalNumber   int64 `json:"total_number,omitempty" xml:"total_number,omitempty"`
}
