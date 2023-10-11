package nrt

// ModifyFundsTypeReqDto 结构体
type ModifyFundsTypeReqDto struct {
	// 分账类型（0卖场，1摊位）
	FundsType int64 `json:"funds_type,omitempty" xml:"funds_type,omitempty"`
	// 阿里摊位id
	StoreId int64 `json:"store_id,omitempty" xml:"store_id,omitempty"`
}
