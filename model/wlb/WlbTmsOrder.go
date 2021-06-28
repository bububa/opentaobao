package wlb

// WlbTmsOrder 
/* model for simplify = false
type WlbTmsOrder struct {

    // 物流订单编号
    
    OrderCode   string `json:"order_code,omitempty"`
    

    // 物流订单的所有者ID,货主
    
    UserId   int64 `json:"user_id,omitempty"`
    

    // 运单号
    
    Code   string `json:"code,omitempty"`
    

    // 物流公司编码
    
    CompanyCode   string `json:"company_code,omitempty"`
    

}
*/

// WlbTmsOrder 
type WlbTmsOrder struct {

    // 物流订单编号
    OrderCode   string `json:"order_code,omitempty"`

    // 物流订单的所有者ID,货主
    UserId   int64 `json:"user_id,omitempty"`

    // 运单号
    Code   string `json:"code,omitempty"`

    // 物流公司编码
    CompanyCode   string `json:"company_code,omitempty"`

}
