package scbp

// AdGroupOperationDto 
/* model for simplify = false
type AdGroupOperationDto struct {

    // 产品id
    
    ProductId   int64 `json:"product_id,omitempty"`
    

    // 线上状态
    
    OnlineStatus   int64 `json:"online_status,omitempty"`
    

    // 产品id
    
    Id   int64 `json:"id,omitempty"`
    

    // key
    
    SettingKey   string `json:"setting_key,omitempty"`
    

    // value
    
    SettingValue   string `json:"setting_value,omitempty"`
    

}
*/

// AdGroupOperationDto 
type AdGroupOperationDto struct {

    // 产品id
    ProductId   int64 `json:"product_id,omitempty"`

    // 线上状态
    OnlineStatus   int64 `json:"online_status,omitempty"`

    // 产品id
    Id   int64 `json:"id,omitempty"`

    // key
    SettingKey   string `json:"setting_key,omitempty"`

    // value
    SettingValue   string `json:"setting_value,omitempty"`

}
