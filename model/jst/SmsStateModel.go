package jst

// SmsStateModel 
/* model for simplify = false
type SmsStateModel struct {

    // 品牌
    
    Manufacturer   string `json:"manufacturer,omitempty"`
    

    // 1 审核中；2 审核不通过；3 待上线；4 已上线；5 已下线
    
    Status   int64 `json:"status,omitempty"`
    

    // success fail
    
    Remark   string `json:"remark,omitempty"`
    

    // CHANNEL（通道）、PUB（LOGO和名称）
    
    Type   string `json:"type,omitempty"`
    

}
*/

// SmsStateModel 
type SmsStateModel struct {

    // 品牌
    Manufacturer   string `json:"manufacturer,omitempty"`

    // 1 审核中；2 审核不通过；3 待上线；4 已上线；5 已下线
    Status   int64 `json:"status,omitempty"`

    // success fail
    Remark   string `json:"remark,omitempty"`

    // CHANNEL（通道）、PUB（LOGO和名称）
    Type   string `json:"type,omitempty"`

}
