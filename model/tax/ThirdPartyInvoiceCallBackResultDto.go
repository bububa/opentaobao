package tax

// ThirdPartyInvoiceCallBackResultDTO 
type ThirdPartyInvoiceCallBackResultDTO struct {
    // 具体明细列表
    ValueList   []ResultItem `json:"value_list,omitempty" xml:"value_list>result_item,omitempty"`
}
