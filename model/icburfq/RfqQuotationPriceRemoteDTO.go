package icburfq

// RfqQuotationPriceRemoteDTO 
type RfqQuotationPriceRemoteDTO struct {
    // 备注
    Remark   string `json:"remark,omitempty" xml:"remark,omitempty"`
    // 预计时间
    EstimatedDate   string `json:"estimated_date,omitempty" xml:"estimated_date,omitempty"`
    // 样品运费支付方
    Payment   string `json:"payment,omitempty" xml:"payment,omitempty"`
    // 是否是免费
    IsFree   string `json:"is_free,omitempty" xml:"is_free,omitempty"`
    // 是否提供样本
    IsSupport   string `json:"is_support,omitempty" xml:"is_support,omitempty"`
}
