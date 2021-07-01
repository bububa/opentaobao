package icburfq

// RfqBuyRequestSearchDetailRemoteDto 结构体
type RfqBuyRequestSearchDetailRemoteDto struct {
	// 附件列表
	Attachments []Attachedfiles `json:"attachments,omitempty" xml:"attachments>attachedfiles,omitempty"`
	// RFQ详情
	RfqDetailDto *BuyRequestSearchDetailRemoteDto `json:"rfq_detail_dto,omitempty" xml:"rfq_detail_dto,omitempty"`
}
