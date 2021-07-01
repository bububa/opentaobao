package axintrade

// AxinFundCreateDto 
type AxinFundCreateDto struct {
    // 支付方支付宝账号id
    PayerAlipayId   string `json:"payer_alipay_id,omitempty" xml:"payer_alipay_id,omitempty"`
    // 支付方账号
    PayerAccount   string `json:"payer_account,omitempty" xml:"payer_account,omitempty"`
    // 支付方昵称
    PayerNick   string `json:"payer_nick,omitempty" xml:"payer_nick,omitempty"`
    // 订单ID
    OuterOrderId   string `json:"outer_order_id,omitempty" xml:"outer_order_id,omitempty"`
    // 付款方式
    PayType   int64 `json:"pay_type,omitempty" xml:"pay_type,omitempty"`
    // 收付款账号类型
    AccountType   int64 `json:"account_type,omitempty" xml:"account_type,omitempty"`
    // 收款方支付宝账号id
    ReceiverAlipayId   string `json:"receiver_alipay_id,omitempty" xml:"receiver_alipay_id,omitempty"`
    // 收款方账号
    ReceiverAccount   string `json:"receiver_account,omitempty" xml:"receiver_account,omitempty"`
    // 收款方昵称
    ReceiverNick   string `json:"receiver_nick,omitempty" xml:"receiver_nick,omitempty"`
    // 二级商户smid
    Smid   string `json:"smid,omitempty" xml:"smid,omitempty"`
    // 正向资金单ID，仅在逆向资金单上使用
    PayFundId   int64 `json:"pay_fund_id,omitempty" xml:"pay_fund_id,omitempty"`
    // 交易方式
    TradeType   string `json:"trade_type,omitempty" xml:"trade_type,omitempty"`
    // 请求版本号，用于幂等校验
    ReqVersion   string `json:"req_version,omitempty" xml:"req_version,omitempty"`
    // 付款来源，默认105直付通
    TradeSource   int64 `json:"trade_source,omitempty" xml:"trade_source,omitempty"`
    // 扩展属性，json格式
    Attributes   string `json:"attributes,omitempty" xml:"attributes,omitempty"`
    // 付款金额总额，单位分
    PayFee   int64 `json:"pay_fee,omitempty" xml:"pay_fee,omitempty"`
    // 支付宝回调地址
    NotifyUrl   string `json:"notify_url,omitempty" xml:"notify_url,omitempty"`
    // 买家实名信息
    ExtUserInfo   *ExtUserInfoDto `json:"ext_user_info,omitempty" xml:"ext_user_info,omitempty"`
    // 商品标题、交易标题、订单关键字等
    Subject   string `json:"subject,omitempty" xml:"subject,omitempty"`
    // 交易渠道,1-借记,2-贷记
    TradeChannel   int64 `json:"trade_channel,omitempty" xml:"trade_channel,omitempty"`
}
