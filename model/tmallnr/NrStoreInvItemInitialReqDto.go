package tmallnr

// NrStoreInvItemInitialReqDto 结构体
type NrStoreInvItemInitialReqDto struct {
	// 如果是品牌商家填写商家的sellerID，如果是零售商需要填写品牌商的sellerID，但是需要平台授权；
	UserId int64 `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// 操作时间
	OperationTime string `json:"operation_time,omitempty" xml:"operation_time,omitempty"`
	// 门店信息
	Stores []StoreDto `json:"stores,omitempty" xml:"stores>store_dto,omitempty"`
	// 操作者
	Operator string `json:"operator,omitempty" xml:"operator,omitempty"`
}
