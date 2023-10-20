package fundplatform

import (
	"sync"
)

// FundAccountJournalResponse 结构体
type FundAccountJournalResponse struct {
	// 流水列表
	JournalList []FundAccountJournalDto `json:"journal_list,omitempty" xml:"journal_list>fund_account_journal_dto,omitempty"`
}

var poolFundAccountJournalResponse = sync.Pool{
	New: func() any {
		return new(FundAccountJournalResponse)
	},
}

// GetFundAccountJournalResponse() 从对象池中获取FundAccountJournalResponse
func GetFundAccountJournalResponse() *FundAccountJournalResponse {
	return poolFundAccountJournalResponse.Get().(*FundAccountJournalResponse)
}

// ReleaseFundAccountJournalResponse 释放FundAccountJournalResponse
func ReleaseFundAccountJournalResponse(v *FundAccountJournalResponse) {
	v.JournalList = v.JournalList[:0]
	poolFundAccountJournalResponse.Put(v)
}
