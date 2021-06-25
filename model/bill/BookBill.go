package bill

// BookBill 
type BookBill struct {

    // 虚拟账户流水编号
    Bid   int64 `json:"bid,omitempty"`

    // 虚拟账户科目编号
    AccountId   int64 `json:"account_id,omitempty"`

    // 流水类型:101-可用金充值;102-可用金扣除;103-冻结;104-解冻;105-冻结金充值;106-冻结金扣除
    JournalType   int64 `json:"journal_type,omitempty"`

    // 操作金额
    Amount   int64 `json:"amount,omitempty"`

    // 记账时间
    BookTime   string `json:"book_time,omitempty"`

    // 备注
    Description   string `json:"description,omitempty"`

    // 创建时间
    GmtCreate   string `json:"gmt_create,omitempty"`

    // 流水的淘宝支付宝id
    TaobaoAlipayId   string `json:"taobao_alipay_id,omitempty"`

    // 流水的商家支付宝id
    OtherAlipayId   string `json:"other_alipay_id,omitempty"`

}
