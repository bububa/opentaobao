package bus

// TopInsProduct 
/* model for simplify = false
type TopInsProduct struct {

    // 保险模块标题
    
    InsTitle   string `json:"ins_title,omitempty"`
    

    // 保险名称
    
    InsName   string `json:"ins_name,omitempty"`
    

    // 保险金额
    
    InsPrice   int64 `json:"ins_price,omitempty"`
    

    // 保险利益点信息
    
    InterestInfo   string `json:"interest_info,omitempty"`
    

    // 保险产品码
    
    ProCode   string `json:"pro_code,omitempty"`
    

    // 资源项，图文信息等
    
    ResourceMap   string `json:"resource_map,omitempty"`
    

}
*/

// TopInsProduct 
type TopInsProduct struct {

    // 保险模块标题
    InsTitle   string `json:"ins_title,omitempty"`

    // 保险名称
    InsName   string `json:"ins_name,omitempty"`

    // 保险金额
    InsPrice   int64 `json:"ins_price,omitempty"`

    // 保险利益点信息
    InterestInfo   string `json:"interest_info,omitempty"`

    // 保险产品码
    ProCode   string `json:"pro_code,omitempty"`

    // 资源项，图文信息等
    ResourceMap   string `json:"resource_map,omitempty"`

}
