package trade

// OrderRefundOperationResponse 结构体
type OrderRefundOperationResponse struct {
	// 不能对某一个订单进行操作时的原因描述
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 不能对某一个订单进行操作时的错误Code
	OperationResultCode string `json:"operation_result_code,omitempty" xml:"operation_result_code,omitempty"`
	// 退款退货操作的Code，由系统定义，目前支持的方式有：refundFeeOnly(仅退款)，refundFeeWithGoods(退货退款),swithGoods(换货)
	OrderRefundActionType string `json:"order_refund_action_type,omitempty" xml:"order_refund_action_type,omitempty"`
	// 操作的用户ID
	OperationUserId string `json:"operation_user_id,omitempty" xml:"operation_user_id,omitempty"`
	// 订单ID
	OrderId string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 是否能对某一个订单进行操作
	CanOperate bool `json:"can_operate,omitempty" xml:"can_operate,omitempty"`
}
