package car

// RentProviderCancelRequest 
type RentProviderCancelRequest struct {
    // 是否确认可以取消
    CancelConfirm   bool `json:"cancel_confirm,omitempty" xml:"cancel_confirm,omitempty"`
    // 服务商ID
    ProviderId   int64 `json:"provider_id,omitempty" xml:"provider_id,omitempty"`
    // 取消拒绝原因
    CancelRejectReason   string `json:"cancel_reject_reason,omitempty" xml:"cancel_reject_reason,omitempty"`
    // 取消拒绝类型
    CancelRejectCode   int64 `json:"cancel_reject_code,omitempty" xml:"cancel_reject_code,omitempty"`
    // 订单id
    OrderId   int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
}
