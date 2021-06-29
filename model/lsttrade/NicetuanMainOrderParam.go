package lsttrade

// NicetuanMainOrderParam 
type NicetuanMainOrderParam struct {
    // 购买用户微信呢称
    NickName   string `json:"nick_name,omitempty" xml:"nick_name,omitempty"`
    // 订单佣金总金额
    RebateAmountTotal   string `json:"rebate_amount_total,omitempty" xml:"rebate_amount_total,omitempty"`
    // 子订单佣金金额
    Recipient   string `json:"recipient,omitempty" xml:"recipient,omitempty"`
    // 父订单号
    OrderCode   string `json:"order_code,omitempty" xml:"order_code,omitempty"`
    // 订单总金额
    AmountTotal   string `json:"amount_total,omitempty" xml:"amount_total,omitempty"`
    // 子单
    NicetuanChildOrderList   []NicetuanChildOrder `json:"nicetuan_child_order_list,omitempty" xml:"nicetuan_child_order_list>nicetuan_child_order,omitempty"`
    // 购买用户微信头像
    Avatar   string `json:"avatar,omitempty" xml:"avatar,omitempty"`
    // 父订单
    NiceTuanPartnerId   string `json:"nice_tuan_partner_id,omitempty" xml:"nice_tuan_partner_id,omitempty"`
    // 团长微信头像
    PartnerAvatar   string `json:"partner_avatar,omitempty" xml:"partner_avatar,omitempty"`
    // 下单日期
    OrderDay   string `json:"order_day,omitempty" xml:"order_day,omitempty"`
}
