package bus

// TaobaoBusSeatpriceGetResultSet 
/* model for simplify = false
type TaobaoBusSeatpriceGetResultSet struct {

    // 错误代码
    
    ErrCode   string `json:"err_code,omitempty"`
    

    // 错误描述
    
    ErrMsg   string `json:"err_msg,omitempty"`
    

    // 余票对象
    
    Module  *struct {
        B2BBusSeatPriceDto  *B2BBusSeatPriceDto `json:"b2b_bus_seat_price_dto,omitempty"`
    } `json:"module,omitempty"`
    

    // serverIP
    
    ServerIP   string `json:"server_i_p,omitempty"`
    

    // 是否查询成功
    
    Success   bool `json:"success,omitempty"`
    

}
*/

// TaobaoBusSeatpriceGetResultSet 
type TaobaoBusSeatpriceGetResultSet struct {

    // 错误代码
    ErrCode   string `json:"err_code,omitempty"`

    // 错误描述
    ErrMsg   string `json:"err_msg,omitempty"`

    // 余票对象
    Module   *B2BBusSeatPriceDto `json:"module,omitempty"`

    // serverIP
    ServerIP   string `json:"server_i_p,omitempty"`

    // 是否查询成功
    Success   bool `json:"success,omitempty"`

}
