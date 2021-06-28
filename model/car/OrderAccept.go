package car

// OrderAccept 
/* model for simplify = false
type OrderAccept struct {

    // 0确认接单 1无法接单
    
    ConfirmType   int64 `json:"confirm_type,omitempty"`
    

    // 拒单原因
    
    Message   string `json:"message,omitempty"`
    

    // 阿里旅行用车订单ID
    
    OrderId   string `json:"order_id,omitempty"`
    

    // 服务商ID,阿里提供
    
    ProviderId   string `json:"provider_id,omitempty"`
    

    // 服务商订单ID,没有可不传
    
    ThirdOrderId   string `json:"third_order_id,omitempty"`
    

    // 可选，卖家id
    
    SellerId   string `json:"seller_id,omitempty"`
    

    // 0:接送机 1：实时打车 2：租车(不传值默认为0)
    
    UseType   int64 `json:"use_type,omitempty"`
    

    // 接单司机纬度
    
    Latitude   string `json:"latitude,omitempty"`
    

    // 接单司机经度
    
    Longitude   string `json:"longitude,omitempty"`
    

    // 接单时间毫秒数
    
    AcceptTime   int64 `json:"accept_time,omitempty"`
    

}
*/

// OrderAccept 
type OrderAccept struct {

    // 0确认接单 1无法接单
    ConfirmType   int64 `json:"confirm_type,omitempty"`

    // 拒单原因
    Message   string `json:"message,omitempty"`

    // 阿里旅行用车订单ID
    OrderId   string `json:"order_id,omitempty"`

    // 服务商ID,阿里提供
    ProviderId   string `json:"provider_id,omitempty"`

    // 服务商订单ID,没有可不传
    ThirdOrderId   string `json:"third_order_id,omitempty"`

    // 可选，卖家id
    SellerId   string `json:"seller_id,omitempty"`

    // 0:接送机 1：实时打车 2：租车(不传值默认为0)
    UseType   int64 `json:"use_type,omitempty"`

    // 接单司机纬度
    Latitude   string `json:"latitude,omitempty"`

    // 接单司机经度
    Longitude   string `json:"longitude,omitempty"`

    // 接单时间毫秒数
    AcceptTime   int64 `json:"accept_time,omitempty"`

}
