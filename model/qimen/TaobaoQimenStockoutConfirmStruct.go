package qimen

// TaobaoQimenStockoutConfirmStruct 
/* model for simplify = false
type TaobaoQimenStockoutConfirmStruct struct {

    // deliveryOrder
    
    DeliveryOrder  *struct {
        DeliveryOrder  *DeliveryOrder `json:"delivery_order,omitempty"`
    } `json:"deliveryOrder,omitempty"`
    

    // packages
    
    Packages  struct {
        Package  []Package `json:"package,omitempty"`
    } `json:"packages,omitempty"`
    

    // orderLines
    
    OrderLines  struct {
        OrderLine  []OrderLine `json:"order_line,omitempty"`
    } `json:"orderLines,omitempty"`
    

    // 扩展属性
    
    ExtendProps  *struct {
        Map  *Map `json:"map,omitempty"`
    } `json:"extendProps,omitempty"`
    

    // 响应结果:success|failure
    
    Flag   string `json:"flag,omitempty"`
    

    // 响应码
    
    Code   string `json:"code,omitempty"`
    

    // 响应信息
    
    Message   string `json:"message,omitempty"`
    

}
*/

// TaobaoQimenStockoutConfirmStruct 
type TaobaoQimenStockoutConfirmStruct struct {

    // deliveryOrder
    DeliveryOrder   *DeliveryOrder `json:"deliveryOrder,omitempty"`

    // packages
    Packages   []Package `json:"packages,omitempty"`

    // orderLines
    OrderLines   []OrderLine `json:"orderLines,omitempty"`

    // 扩展属性
    ExtendProps   *Map `json:"extendProps,omitempty"`

    // 响应结果:success|failure
    Flag   string `json:"flag,omitempty"`

    // 响应码
    Code   string `json:"code,omitempty"`

    // 响应信息
    Message   string `json:"message,omitempty"`

}
