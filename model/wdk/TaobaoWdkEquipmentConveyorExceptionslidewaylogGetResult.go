package wdk

// TaobaoWdkEquipmentConveyorExceptionslidewaylogGetResult 
/* model for simplify = false
type TaobaoWdkEquipmentConveyorExceptionslidewaylogGetResult struct {

    // 返回值与返回的信息
    
    ResultCode  *struct {
        ResultCode  *ResultCode `json:"result_code,omitempty"`
    } `json:"result_code,omitempty"`
    

    // 返回的数据
    
    Data   string `json:"data,omitempty"`
    

}
*/

// TaobaoWdkEquipmentConveyorExceptionslidewaylogGetResult 
type TaobaoWdkEquipmentConveyorExceptionslidewaylogGetResult struct {

    // 返回值与返回的信息
    ResultCode   *ResultCode `json:"result_code,omitempty"`

    // 返回的数据
    Data   string `json:"data,omitempty"`

}
