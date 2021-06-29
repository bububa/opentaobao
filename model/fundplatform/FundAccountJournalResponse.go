package fundplatform

// FundAccountJournalResponse 
type FundAccountJournalResponse struct {
    // 流水列表
    JournalList   []FundAccountJournalDTO `json:"journal_list,omitempty" xml:"journal_list>fund_account_journal_dto,omitempty"`
}
