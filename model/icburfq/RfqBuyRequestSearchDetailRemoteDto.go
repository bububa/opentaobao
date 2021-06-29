package icburfq

// RfqBuyRequestSearchDetailRemoteDTO 
type RfqBuyRequestSearchDetailRemoteDTO struct {
    // 附件列表
    Attachments   []Attachedfiles `json:"attachments,omitempty" xml:"attachments>attachedfiles,omitempty"`
    // RFQ详情
    RfqDetailDto   *BuyRequestSearchDetailRemoteDTO `json:"rfq_detail_dto,omitempty" xml:"rfq_detail_dto,omitempty"`
}
