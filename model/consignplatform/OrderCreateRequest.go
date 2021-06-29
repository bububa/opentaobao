package consignplatform

// OrderCreateRequest 
type OrderCreateRequest struct {
    // 商家备注
    UserMemo   string `json:"user_memo,omitempty" xml:"user_memo,omitempty"`
    // 收件地址
    ReceiveAddress   *AddressDtoForTop `json:"receive_address,omitempty" xml:"receive_address,omitempty"`
    // 收件人
    Receiver   *PersonDTO `json:"receiver,omitempty" xml:"receiver,omitempty"`
    // 买家留言
    BuyerMemo   string `json:"buyer_memo,omitempty" xml:"buyer_memo,omitempty"`
    // 外部订单id，会用来做幂等，需要保证唯一
    OuterOrderId   string `json:"outer_order_id,omitempty" xml:"outer_order_id,omitempty"`
    // 订单来源
    OrderSource   string `json:"order_source,omitempty" xml:"order_source,omitempty"`
    // 子订单列表
    SubOrderList   []SubOrderDTO `json:"sub_order_list,omitempty" xml:"sub_order_list>sub_order_dto,omitempty"`
}
