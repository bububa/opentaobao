package aecreatives

// AffiliateProductResultDto 结构体
type AffiliateProductResultDto struct {
	// 商品列表
	Products []Product `json:"products,omitempty" xml:"products>product,omitempty"`
	// 当前页记录数
	CurrentRecordCount int64 `json:"current_record_count,omitempty" xml:"current_record_count,omitempty"`
}
