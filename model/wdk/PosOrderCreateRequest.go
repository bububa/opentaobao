package wdk

// PosOrderCreateRequest 
/* model for simplify = false
type PosOrderCreateRequest struct {

    // 子订单列表
    
    SubOrderDOList  struct {
        PosSubOrderDo  []PosSubOrderDo `json:"pos_sub_order_do,omitempty"`
    } `json:"sub_order_d_o_list,omitempty"`
    

    // 支付时间，必填
    
    PayTime   string `json:"pay_time,omitempty"`
    

    // 外部主订单号，必填
    
    OutOrderId   string `json:"out_order_id,omitempty"`
    

    // 经营店code，必填
    
    StoreId   string `json:"store_id,omitempty"`
    

    // 渠道店id
    
    ShopId   string `json:"shop_id,omitempty"`
    

    // 会员卡号
    
    MemberCardNum   string `json:"member_card_num,omitempty"`
    

    // 支付方式
    
    PayChannelList  struct {
        PosPayChannel  []PosPayChannel `json:"pos_pay_channel,omitempty"`
    } `json:"pay_channel_list,omitempty"`
    

    // 兼容老接口的数据
    
    OldData   int64 `json:"old_data,omitempty"`
    

}
*/

// PosOrderCreateRequest 
type PosOrderCreateRequest struct {

    // 子订单列表
    SubOrderDOList   []PosSubOrderDo `json:"sub_order_d_o_list,omitempty"`

    // 支付时间，必填
    PayTime   string `json:"pay_time,omitempty"`

    // 外部主订单号，必填
    OutOrderId   string `json:"out_order_id,omitempty"`

    // 经营店code，必填
    StoreId   string `json:"store_id,omitempty"`

    // 渠道店id
    ShopId   string `json:"shop_id,omitempty"`

    // 会员卡号
    MemberCardNum   string `json:"member_card_num,omitempty"`

    // 支付方式
    PayChannelList   []PosPayChannel `json:"pay_channel_list,omitempty"`

    // 兼容老接口的数据
    OldData   int64 `json:"old_data,omitempty"`

}
