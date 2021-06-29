package alicom

// ExchangeModel 
type ExchangeModel struct {
    // 从tae获取的混淆nick
    Nick   string `json:"nick,omitempty" xml:"nick,omitempty"`
    // 商家订单编号
    OrderNo   string `json:"order_no,omitempty" xml:"order_no,omitempty"`
    // 兑换优惠券的金额，单位：分
    Amount   string `json:"amount,omitempty" xml:"amount,omitempty"`
    // 扩展信息
    Ext   *Ext `json:"ext,omitempty" xml:"ext,omitempty"`
}
