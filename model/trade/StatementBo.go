package trade

// StatementBo 
/* model for simplify = false
type StatementBo struct {

    // 订单金额
    
    Amount   int64 `json:"amount,omitempty"`
    

    // 对账单号
    
    BillNo   string `json:"bill_no,omitempty"`
    

    // 渠道编码 WX_APPLET("10001", "微信小程序"),     WX_OFFICIAL_SERVER("10002", "微信服务号/微信商城"),     ALIPAY_APPLET("20001", "支付宝小程序"),     OFFLINE_MEMBER_CARD("30001", "线下会员卡"),     MERCHANT_APP("40001", "商家自有app"),     TXD_APP("50001", "淘鲜达app"),     TXD_SELF_POS("60001", "淘鲜达自助pos"),     TXD_ARTIFACT_POS("60001", "淘鲜达人工pos");
    
    ChannelCode   string `json:"channel_code,omitempty"`
    

    // 订单创建时间
    
    CreateTime   string `json:"create_time,omitempty"`
    

    // 外部会员Id
    
    OpenMemberId   string `json:"open_member_id,omitempty"`
    

    // 渠道订单Id
    
    OrderId   string `json:"order_id,omitempty"`
    

    // 订单类型     NORMAL(0, "正向"),     REFUND(1, "逆向"),     REVOKED(2, "撤销");
    
    OrderType   int64 `json:"order_type,omitempty"`
    

    // 支付方式 ALIPAY(1, "支付宝"),     WECHAT(2, "微信"),     CASH(3, "现金"),     VALUE_CARD(4, "储值卡"),     OTHER(99, "其他")
    
    PayType   int64 `json:"pay_type,omitempty"`
    

}
*/

// StatementBo 
type StatementBo struct {

    // 订单金额
    Amount   int64 `json:"amount,omitempty"`

    // 对账单号
    BillNo   string `json:"bill_no,omitempty"`

    // 渠道编码 WX_APPLET("10001", "微信小程序"),     WX_OFFICIAL_SERVER("10002", "微信服务号/微信商城"),     ALIPAY_APPLET("20001", "支付宝小程序"),     OFFLINE_MEMBER_CARD("30001", "线下会员卡"),     MERCHANT_APP("40001", "商家自有app"),     TXD_APP("50001", "淘鲜达app"),     TXD_SELF_POS("60001", "淘鲜达自助pos"),     TXD_ARTIFACT_POS("60001", "淘鲜达人工pos");
    ChannelCode   string `json:"channel_code,omitempty"`

    // 订单创建时间
    CreateTime   string `json:"create_time,omitempty"`

    // 外部会员Id
    OpenMemberId   string `json:"open_member_id,omitempty"`

    // 渠道订单Id
    OrderId   string `json:"order_id,omitempty"`

    // 订单类型     NORMAL(0, "正向"),     REFUND(1, "逆向"),     REVOKED(2, "撤销");
    OrderType   int64 `json:"order_type,omitempty"`

    // 支付方式 ALIPAY(1, "支付宝"),     WECHAT(2, "微信"),     CASH(3, "现金"),     VALUE_CARD(4, "储值卡"),     OTHER(99, "其他")
    PayType   int64 `json:"pay_type,omitempty"`

}
