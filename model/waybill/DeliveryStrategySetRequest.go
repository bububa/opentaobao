package waybill

// DeliveryStrategySetRequest 
type DeliveryStrategySetRequest struct {
    // 策略信息对象
    DeliveryStrategyInfo   *DeliveryStrategyInfo `json:"delivery_strategy_info,omitempty" xml:"delivery_strategy_info,omitempty"`
}
