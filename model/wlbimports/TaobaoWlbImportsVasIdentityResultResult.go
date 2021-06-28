package wlbimports

// TaobaoWlbImportsVasIdentityResultResult 
/* model for simplify = false
type TaobaoWlbImportsVasIdentityResultResult struct {

    // 物流订单号
    
    LgOrderCode   string `json:"lg_order_code,omitempty"`
    

    // 鉴定结果项
    
    VasResults  struct {
        IdentityItemDto  []IdentityItemDto `json:"identity_item_dto,omitempty"`
    } `json:"vas_results,omitempty"`
    

}
*/

// TaobaoWlbImportsVasIdentityResultResult 
type TaobaoWlbImportsVasIdentityResultResult struct {

    // 物流订单号
    LgOrderCode   string `json:"lg_order_code,omitempty"`

    // 鉴定结果项
    VasResults   []IdentityItemDto `json:"vas_results,omitempty"`

}
