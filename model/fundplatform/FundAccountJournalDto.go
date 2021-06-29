package fundplatform

// FundAccountJournalDto 
type FundAccountJournalDto struct {

    // 账户ID
    
    AccountId   int64 `json:"account_id,omitempty" xml:"account_id,omitempty"`
    

    // 操作金额
    
    DealAmount   int64 `json:"deal_amount,omitempty" xml:"deal_amount,omitempty"`
    

    // 描述
    
    Description   string `json:"description,omitempty" xml:"description,omitempty"`
    

    // 创建时间
    
    GmtCreate   string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
    

    // 最后修改时间
    
    GmtModified   string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
    

    // 流水类型：FUND_ACCOUNT_IN 充值 FUND_ACCOUNT_OUT 提现 FUND_ACCOUNT_DEDUCT 扣款 FUND_ACCOUNT_FREEZE 冻结 FUND_ACCOUNT_UNFREEZE 解冻
    
    JournalType   string `json:"journal_type,omitempty" xml:"journal_type,omitempty"`
    

    // 外部订单号
    
    OutBizId   string `json:"out_biz_id,omitempty" xml:"out_biz_id,omitempty"`
    

    // 用户ID
    
    UserId   int64 `json:"user_id,omitempty" xml:"user_id,omitempty"`
    

}
