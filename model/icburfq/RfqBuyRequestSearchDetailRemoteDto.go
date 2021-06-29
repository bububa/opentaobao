package icburfq

// RfqBuyRequestSearchDetailRemoteDto 
type RfqBuyRequestSearchDetailRemoteDto struct {

    // 附件列表
    
    Attachments   []Attachedfiles `json:"attachments,omitempty" xml:"attachments,omitempty"`
    

    // RFQ详情
    
    RfqDetailDto   *BuyRequestSearchDetailRemoteDto `json:"rfq_detail_dto,omitempty" xml:"rfq_detail_dto,omitempty"`
    

}
