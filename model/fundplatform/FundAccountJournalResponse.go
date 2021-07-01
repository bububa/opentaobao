package fundplatform

// FundAccountJournalResponse 结构体
type FundAccountJournalResponse struct {
	// 流水列表
	JournalList []FundAccountJournalDto `json:"journal_list,omitempty" xml:"journal_list>fund_account_journal_dto,omitempty"`
}
