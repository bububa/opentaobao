package fundplatform

// FundAccountJournalQueryReq 结构体
type FundAccountJournalQueryReq struct {
	// 账户ID
	AccountId int64 `json:"account_id,omitempty" xml:"account_id,omitempty"`
	// 流水创建开始时间，前闭后开
	CreateBeginTime string `json:"create_begin_time,omitempty" xml:"create_begin_time,omitempty"`
	// 流水创建结束时间，前闭后开
	CreateEndTime string `json:"create_end_time,omitempty" xml:"create_end_time,omitempty"`
	// 当前页，从1开始
	CurrentPageNo int64 `json:"current_page_no,omitempty" xml:"current_page_no,omitempty"`
	// 查询流水类型枚举，为空代表所有类型FUND_ACCOUNT_IN 充值     FUND_ACCOUNT_OUT 提现     FUND_ACCOUNT_DEDUCT 扣款     FUND_ACCOUNT_FREEZE 冻结     FUND_ACCOUNT_UNFREEZE 解冻
	JournalTypes []string `json:"journal_types,omitempty" xml:"journal_types>string,omitempty"`
	// 页大小，[1,1000]
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
}
