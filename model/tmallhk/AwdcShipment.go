package tmallhk

// AwdcShipment 
/* model for simplify = false
type AwdcShipment struct {

    // 到达城市
    
    ArrivalCity   string `json:"arrival_city,omitempty"`
    

    // 抵港时间
    
    ArrivalDate   string `json:"arrival_date,omitempty"`
    

    // 起运城市
    
    DepartureCity   string `json:"departure_city,omitempty"`
    

    // 离港时间
    
    DepartureDate   string `json:"departure_date,omitempty"`
    

    // 报关开始时间
    
    DoIn   string `json:"do_in,omitempty"`
    

    // 报关单号
    
    DoNumber   string `json:"do_number,omitempty"`
    

    // 报关结束时间
    
    DoOut   string `json:"do_out,omitempty"`
    

    // 押运公司物流单号
    
    LogisticNumber   string `json:"logistic_number,omitempty"`
    

    // 航班装运订单
    
    ShipmentNumber   string `json:"shipment_number,omitempty"`
    

    // 押运公司
    
    Shipper   string `json:"shipper,omitempty"`
    

}
*/

// AwdcShipment 
type AwdcShipment struct {

    // 到达城市
    ArrivalCity   string `json:"arrival_city,omitempty"`

    // 抵港时间
    ArrivalDate   string `json:"arrival_date,omitempty"`

    // 起运城市
    DepartureCity   string `json:"departure_city,omitempty"`

    // 离港时间
    DepartureDate   string `json:"departure_date,omitempty"`

    // 报关开始时间
    DoIn   string `json:"do_in,omitempty"`

    // 报关单号
    DoNumber   string `json:"do_number,omitempty"`

    // 报关结束时间
    DoOut   string `json:"do_out,omitempty"`

    // 押运公司物流单号
    LogisticNumber   string `json:"logistic_number,omitempty"`

    // 航班装运订单
    ShipmentNumber   string `json:"shipment_number,omitempty"`

    // 押运公司
    Shipper   string `json:"shipper,omitempty"`

}
