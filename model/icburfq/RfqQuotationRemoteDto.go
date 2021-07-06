package icburfq

// RfqQuotationRemoteDto 结构体
type RfqQuotationRemoteDto struct {
	// 报价列表
	PriceList []PriceList `json:"price_list,omitempty" xml:"price_list>price_list,omitempty"`
	// 给买家留言
	Details string `json:"details,omitempty" xml:"details,omitempty"`
	// 附件file_str,请通过调用alibaba.icbu.annex.upload结果作为入参
	AnnexFilesStr string `json:"annex_files_str,omitempty" xml:"annex_files_str,omitempty"`
	// RFQID
	RfqId string `json:"rfq_id,omitempty" xml:"rfq_id,omitempty"`
	// 付费条款
	PaymentTerms string `json:"payment_terms,omitempty" xml:"payment_terms,omitempty"`
	// 过期时间
	ExpiryDate string `json:"expiry_date,omitempty" xml:"expiry_date,omitempty"`
	// 样本
	Sample *RfqQuotationPriceRemoteDto `json:"sample,omitempty" xml:"sample,omitempty"`
	// 报价ID
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}
