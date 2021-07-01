package waybill

// DeliveryStrategySetRequest 结构体
type DeliveryStrategySetRequest struct {
	// 策略信息对象
	DeliveryStrategyInfo *DeliveryStrategyInfo `json:"delivery_strategy_info,omitempty" xml:"delivery_strategy_info,omitempty"`
}
