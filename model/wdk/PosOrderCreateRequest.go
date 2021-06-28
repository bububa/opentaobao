package wdk

// PosOrderCreateRequest 
type PosOrderCreateRequest struct {

    // 子订单列表
    
    SubOrderDOList   []PosSubOrderDo `json:"sub_order_d_o_list,omitempty" xml:"sub_order_d_o_list,omitempty"`
    

    // 支付时间，必填
    
    PayTime   string `json:"pay_time,omitempty" xml:"pay_time,omitempty"`
    

    // 外部主订单号，必填
    
    OutOrderId   string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
    

    // 经营店code，必填
    
    StoreId   string `json:"store_id,omitempty" xml:"store_id,omitempty"`
    

    // 渠道店id
    
    ShopId   string `json:"shop_id,omitempty" xml:"shop_id,omitempty"`
    

    // 会员卡号
    
    MemberCardNum   string `json:"member_card_num,omitempty" xml:"member_card_num,omitempty"`
    

    // 支付方式
    
    PayChannelList   []PosPayChannel `json:"pay_channel_list,omitempty" xml:"pay_channel_list,omitempty"`
    

    // 兼容老接口的数据
    
    OldData   int64 `json:"old_data,omitempty" xml:"old_data,omitempty"`
    

}
