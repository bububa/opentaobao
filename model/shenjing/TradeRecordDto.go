package shenjing

// TradeRecordDto 
type TradeRecordDto struct {

    // 账户币种
    
    CurrencyType   string `json:"currency_type,omitempty" xml:"currency_type,omitempty"`
    

    // 摘要
    
    Summary   string `json:"summary,omitempty" xml:"summary,omitempty"`
    

    // 唯一编号（网关流水唯一键）
    
    UniqueNo   string `json:"unique_no,omitempty" xml:"unique_no,omitempty"`
    

    // 银行账号编码
    
    BankCode   string `json:"bank_code,omitempty" xml:"bank_code,omitempty"`
    

    // 业务参考号
    
    BizRefNo   string `json:"biz_ref_no,omitempty" xml:"biz_ref_no,omitempty"`
    

    // 对方账户
    
    OtherSideAccountNo   string `json:"other_side_account_no,omitempty" xml:"other_side_account_no,omitempty"`
    

    // 银行交易流水号
    
    TradeNo   string `json:"trade_no,omitempty" xml:"trade_no,omitempty"`
    

    // 付款目的
    
    Purpose   string `json:"purpose,omitempty" xml:"purpose,omitempty"`
    

    // 流水同步日期
    
    SyncDate   int64 `json:"sync_date,omitempty" xml:"sync_date,omitempty"`
    

    // 交易日期
    
    TradeDate   int64 `json:"trade_date,omitempty" xml:"trade_date,omitempty"`
    

    // 借方金额(支出)
    
    DebitAmount   string `json:"debit_amount,omitempty" xml:"debit_amount,omitempty"`
    

    // 业务摘要
    
    BizSummary   string `json:"biz_summary,omitempty" xml:"biz_summary,omitempty"`
    

    // 流程实例号（只有招商银行才有）
    
    ProcessNo   string `json:"process_no,omitempty" xml:"process_no,omitempty"`
    

    // 账号当前余额
    
    Balance   string `json:"balance,omitempty" xml:"balance,omitempty"`
    

    // 其他摘要
    
    OtherSummary   string `json:"other_summary,omitempty" xml:"other_summary,omitempty"`
    

    // 对方账户名称
    
    OtherSideAccountName   string `json:"other_side_account_name,omitempty" xml:"other_side_account_name,omitempty"`
    

    // 业务名称
    
    BizName   string `json:"biz_name,omitempty" xml:"biz_name,omitempty"`
    

    // 账户号码
    
    AccountNo   string `json:"account_no,omitempty" xml:"account_no,omitempty"`
    

    // 贷方金额(收入)
    
    CreditAmount   string `json:"credit_amount,omitempty" xml:"credit_amount,omitempty"`
    

    // 专属账号
    
    AccountNoEx   string `json:"account_no_ex,omitempty" xml:"account_no_ex,omitempty"`
    

    // 交易类型
    
    TradeType   string `json:"trade_type,omitempty" xml:"trade_type,omitempty"`
    

}
