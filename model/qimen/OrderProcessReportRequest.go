package qimen

// OrderProcessReportRequest 
/* model for simplify = false
type OrderProcessReportRequest struct {

    // 订单信息
    
    Order  *struct {
        Order  *Order `json:"order,omitempty"`
    } `json:"order,omitempty"`
    

    // 订单处理信息
    
    Process  *struct {
        Process  *Process `json:"process,omitempty"`
    } `json:"process,omitempty"`
    

    // 扩展属性
    
    ExtendProps  *struct {
        Map  *Map `json:"map,omitempty"`
    } `json:"extendProps,omitempty"`
    

    // 备注
    
    Remark   string `json:"remark,omitempty"`
    

}
*/

// OrderProcessReportRequest 
type OrderProcessReportRequest struct {

    // 订单信息
    Order   *Order `json:"order,omitempty"`

    // 订单处理信息
    Process   *Process `json:"process,omitempty"`

    // 扩展属性
    ExtendProps   *Map `json:"extendProps,omitempty"`

    // 备注
    Remark   string `json:"remark,omitempty"`

}
