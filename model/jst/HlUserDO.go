package jst

// HlUserDO 
/* model for simplify = false
type HlUserDO struct {

    // 如果为空，则默认是X_TO_SYSTEM,X_WAIT_ALLOCATION,X_OUT_WAREHOUSE
    
    OpenNodes   string `json:"open_nodes,omitempty"`
    

    // 回流信息是否开通买家端展示
    
    OpenForBuyer   string `json:"open_for_buyer,omitempty"`
    

}
*/

// HlUserDO 
type HlUserDO struct {

    // 如果为空，则默认是X_TO_SYSTEM,X_WAIT_ALLOCATION,X_OUT_WAREHOUSE
    OpenNodes   string `json:"open_nodes,omitempty"`

    // 回流信息是否开通买家端展示
    OpenForBuyer   string `json:"open_for_buyer,omitempty"`

}
