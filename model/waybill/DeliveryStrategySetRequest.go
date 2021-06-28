package waybill

// DeliveryStrategySetRequest 
/* model for simplify = false
type DeliveryStrategySetRequest struct {

    // 策略信息对象
    
    DeliveryStrategyInfo  *struct {
        DeliveryStrategyInfo  *DeliveryStrategyInfo `json:"delivery_strategy_info,omitempty"`
    } `json:"delivery_strategy_info,omitempty"`
    

}
*/

// DeliveryStrategySetRequest 
type DeliveryStrategySetRequest struct {

    // 策略信息对象
    DeliveryStrategyInfo   *DeliveryStrategyInfo `json:"delivery_strategy_info,omitempty"`

}
