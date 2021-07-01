package mos

// PosBankSaleInfoDto 
type PosBankSaleInfoDto struct {
    // 订单号，唯一幂等字段
    OrderId   string `json:"order_id,omitempty" xml:"order_id,omitempty"`
    // 原订单号
    OriginalOrderId   string `json:"original_order_id,omitempty" xml:"original_order_id,omitempty"`
    // 小票号
    SequenceNo   string `json:"sequence_no,omitempty" xml:"sequence_no,omitempty"`
    // 支付渠道。001：云闪付、002：银行卡
    PaymentChannel   string `json:"payment_channel,omitempty" xml:"payment_channel,omitempty"`
    // 卡号
    CardNo   string `json:"card_no,omitempty" xml:"card_no,omitempty"`
    // 交易时间。yyyy-MM-dd HH:mm:ss
    TradeTime   string `json:"trade_time,omitempty" xml:"trade_time,omitempty"`
    // 销售类型。0：销售，1：退款退货
    SaleType   string `json:"sale_type,omitempty" xml:"sale_type,omitempty"`
    // 销售类型子类。001：消费，002：撤销，006：退款退货
    SubSaleType   string `json:"sub_sale_type,omitempty" xml:"sub_sale_type,omitempty"`
    // 交易金额。单位：分
    TradeAmount   int64 `json:"trade_amount,omitempty" xml:"trade_amount,omitempty"`
    // 银行编码
    BankCode   string `json:"bank_code,omitempty" xml:"bank_code,omitempty"`
    // 银行名字
    BankName   string `json:"bank_name,omitempty" xml:"bank_name,omitempty"`
    // 终端编号
    TerminalNo   string `json:"terminal_no,omitempty" xml:"terminal_no,omitempty"`
    // 银行商户号
    BankShopNo   string `json:"bank_shop_no,omitempty" xml:"bank_shop_no,omitempty"`
    // 商户流水号，查询流水号
    PosTraceNo   string `json:"pos_trace_no,omitempty" xml:"pos_trace_no,omitempty"`
    // 批次号
    BatchNo   string `json:"batch_no,omitempty" xml:"batch_no,omitempty"`
    // 交易参考号，系统跟踪号
    TradeRefNo   string `json:"trade_ref_no,omitempty" xml:"trade_ref_no,omitempty"`
    // 支付类型行号
    PayTypeNo   int64 `json:"pay_type_no,omitempty" xml:"pay_type_no,omitempty"`
    // 外部门店号，类似HZ01
    StoreNo   string `json:"store_no,omitempty" xml:"store_no,omitempty"`
    // 操作时间。yyyy-MM-dd HH:mm:ss
    OperateTime   string `json:"operate_time,omitempty" xml:"operate_time,omitempty"`
    // 冲正标志，供银行卡调账使用，表字段需要，不由接口传入，默认为00
    Rback   string `json:"rback,omitempty" xml:"rback,omitempty"`
    // 备注
    Remark   string `json:"remark,omitempty" xml:"remark,omitempty"`
    // 收银员
    Cashier   string `json:"cashier,omitempty" xml:"cashier,omitempty"`
    // ip地址
    Ip   string `json:"ip,omitempty" xml:"ip,omitempty"`
    // mac地址
    Mac   string `json:"mac,omitempty" xml:"mac,omitempty"`
    // 通讯类型。001：网线，002:电话
    CommunicateType   string `json:"communicate_type,omitempty" xml:"communicate_type,omitempty"`
    // 加密摘要，未脱敏号加密
    EncrypteSummary   string `json:"encrypte_summary,omitempty" xml:"encrypte_summary,omitempty"`
    // 扩展字段
    ExtendParam   string `json:"extend_param,omitempty" xml:"extend_param,omitempty"`
}
