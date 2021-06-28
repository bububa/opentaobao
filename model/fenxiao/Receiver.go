package fenxiao

// Receiver 
/* model for simplify = false
type Receiver struct {

    // 收货人全名
    
    Name   string `json:"name,omitempty"`
    

    // 固定电话
    
    Phone   string `json:"phone,omitempty"`
    

    // 移动电话
    
    MobilePhone   string `json:"mobile_phone,omitempty"`
    

    // 收货人的详细地址信息
    
    Address   string `json:"address,omitempty"`
    

    // 区/县
    
    District   string `json:"district,omitempty"`
    

    // 收货人的城市
    
    City   string `json:"city,omitempty"`
    

    // 邮政编码
    
    Zip   string `json:"zip,omitempty"`
    

    // 收货人所在省份
    
    State   string `json:"state,omitempty"`
    

    // 证件号
    
    CardId   string `json:"card_id,omitempty"`
    

}
*/

// Receiver 
type Receiver struct {

    // 收货人全名
    Name   string `json:"name,omitempty"`

    // 固定电话
    Phone   string `json:"phone,omitempty"`

    // 移动电话
    MobilePhone   string `json:"mobile_phone,omitempty"`

    // 收货人的详细地址信息
    Address   string `json:"address,omitempty"`

    // 区/县
    District   string `json:"district,omitempty"`

    // 收货人的城市
    City   string `json:"city,omitempty"`

    // 邮政编码
    Zip   string `json:"zip,omitempty"`

    // 收货人所在省份
    State   string `json:"state,omitempty"`

    // 证件号
    CardId   string `json:"card_id,omitempty"`

}
