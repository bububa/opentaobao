package mos

// PosSaleOrderDto 
/* model for simplify = false
type PosSaleOrderDto struct {

    // 扩展信息
    
    ExtendParams   string `json:"extend_params,omitempty"`
    

    // 会员卡号
    
    MemberCardNo   string `json:"member_card_no,omitempty"`
    

    // 会员电话
    
    MemberMobile   string `json:"member_mobile,omitempty"`
    

    // 操作员
    
    Operator   string `json:"operator,omitempty"`
    

    // 商品列表
    
    SaleItems  struct {
        PosOrderSaleItemDto  []PosOrderSaleItemDto `json:"pos_order_sale_item_dto,omitempty"`
    } `json:"sale_items,omitempty"`
    

    // 开票单号
    
    SaleTicketNo   string `json:"sale_ticket_no,omitempty"`
    

}
*/

// PosSaleOrderDto 
type PosSaleOrderDto struct {

    // 扩展信息
    ExtendParams   string `json:"extend_params,omitempty"`

    // 会员卡号
    MemberCardNo   string `json:"member_card_no,omitempty"`

    // 会员电话
    MemberMobile   string `json:"member_mobile,omitempty"`

    // 操作员
    Operator   string `json:"operator,omitempty"`

    // 商品列表
    SaleItems   []PosOrderSaleItemDto `json:"sale_items,omitempty"`

    // 开票单号
    SaleTicketNo   string `json:"sale_ticket_no,omitempty"`

}
