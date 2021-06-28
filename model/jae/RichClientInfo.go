package jae

// RichClientInfo 
/* model for simplify = false
type RichClientInfo struct {

    // appkey
    
    Appkey   string `json:"appkey,omitempty"`
    

    // 用户昵称
    
    BuyerNick   string `json:"buyer_nick,omitempty"`
    

    // 经度
    
    Lng   string `json:"lng,omitempty"`
    

    // 用户id
    
    BuyerId   int64 `json:"buyer_id,omitempty"`
    

    // 纬度
    
    Lat   string `json:"lat,omitempty"`
    

    // 设备号
    
    DeviceId   string `json:"device_id,omitempty"`
    

    // 来源ip
    
    Ip   string `json:"ip,omitempty"`
    

}
*/

// RichClientInfo 
type RichClientInfo struct {

    // appkey
    Appkey   string `json:"appkey,omitempty"`

    // 用户昵称
    BuyerNick   string `json:"buyer_nick,omitempty"`

    // 经度
    Lng   string `json:"lng,omitempty"`

    // 用户id
    BuyerId   int64 `json:"buyer_id,omitempty"`

    // 纬度
    Lat   string `json:"lat,omitempty"`

    // 设备号
    DeviceId   string `json:"device_id,omitempty"`

    // 来源ip
    Ip   string `json:"ip,omitempty"`

}
