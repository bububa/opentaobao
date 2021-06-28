package mos

// PosSaleOrderDto 
type PosSaleOrderDto struct {

    // 扩展信息
    
    ExtendParams   string `json:"extend_params,omitempty" xml:"extend_params,omitempty"`
    

    // 会员卡号
    
    MemberCardNo   string `json:"member_card_no,omitempty" xml:"member_card_no,omitempty"`
    

    // 会员电话
    
    MemberMobile   string `json:"member_mobile,omitempty" xml:"member_mobile,omitempty"`
    

    // 操作员
    
    Operator   string `json:"operator,omitempty" xml:"operator,omitempty"`
    

    // 商品列表
    
    SaleItems   []PosOrderSaleItemDto `json:"sale_items,omitempty" xml:"sale_items,omitempty"`
    

    // 开票单号
    
    SaleTicketNo   string `json:"sale_ticket_no,omitempty" xml:"sale_ticket_no,omitempty"`
    

}
