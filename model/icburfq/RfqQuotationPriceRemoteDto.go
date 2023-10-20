package icburfq

// RfqQuotationPriceRemoteDto 结构体
type RfqQuotationPriceRemoteDto struct {
	// 备注
	Remark string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 样品运费支付方
	Payment string `json:"payment,omitempty" xml:"payment,omitempty"`
	// 是否是免费
	IsFree string `json:"is_free,omitempty" xml:"is_free,omitempty"`
	// 是否提供样本
	IsSupport string `json:"is_support,omitempty" xml:"is_support,omitempty"`
	// 预计时间
	EstimatedDate float64 `json:"estimated_date,omitempty" xml:"estimated_date,omitempty"`
}
