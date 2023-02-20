package aecreatives

type AffiliateProductResultDto struct {
	// 返回商品列表
	Products []Product `json:"products,omitempty" xml:"products>product,omitempty"`
	// 当前返回数量
	CurrentRecordCount int64 `json:"current_record_count,omitempty" xml:"current_record_count,omitempty"`
}
