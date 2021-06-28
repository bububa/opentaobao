package logistic

// ReachableAddressAndServiceDto 
/* model for simplify = false
type ReachableAddressAndServiceDto struct {

    // 收货地址
    
    ReceiveAddress  *struct {
        ReceiveAddress  *ReceiveAddress `json:"receive_address,omitempty"`
    } `json:"receive_address,omitempty"`
    

    // 服务列表,每一项必须为json的string格式,快运必填,快递为空则默认为'标准快递'
    
    ServiceCodeList  struct {
        String  []string `json:"string,omitempty"`
    } `json:"service_code_list,omitempty"`
    

    // 每条收发地址的key,用户自定义,每次请求多个地址不能重复
    
    ObjectId   string `json:"object_id,omitempty"`
    

    // 发货地址
    
    SendAddress  *struct {
        AddressDto  *AddressDto `json:"address_dto,omitempty"`
    } `json:"send_address,omitempty"`
    

    // 淘宝开放地址ID
    
    Oaid   string `json:"oaid,omitempty"`
    

    // 订单id
    
    OrderId   int64 `json:"order_id,omitempty"`
    

    // C2M&1688开放地址ID
    
    Caid   string `json:"caid,omitempty"`
    

}
*/

// ReachableAddressAndServiceDto 
type ReachableAddressAndServiceDto struct {

    // 收货地址
    ReceiveAddress   *ReceiveAddress `json:"receive_address,omitempty"`

    // 服务列表,每一项必须为json的string格式,快运必填,快递为空则默认为'标准快递'
    ServiceCodeList   []string `json:"service_code_list,omitempty"`

    // 每条收发地址的key,用户自定义,每次请求多个地址不能重复
    ObjectId   string `json:"object_id,omitempty"`

    // 发货地址
    SendAddress   *AddressDto `json:"send_address,omitempty"`

    // 淘宝开放地址ID
    Oaid   string `json:"oaid,omitempty"`

    // 订单id
    OrderId   int64 `json:"order_id,omitempty"`

    // C2M&1688开放地址ID
    Caid   string `json:"caid,omitempty"`

}
