package wlbimports

// TaobaoWlbImportsVasIdentityResultResult 
type TaobaoWlbImportsVasIdentityResultResult struct {
    // 物流订单号
    LgOrderCode   string `json:"lg_order_code,omitempty" xml:"lg_order_code,omitempty"`
    // 鉴定结果项
    VasResults   []IdentityItemDto `json:"vas_results,omitempty" xml:"vas_results>identity_item_dto,omitempty"`
}
