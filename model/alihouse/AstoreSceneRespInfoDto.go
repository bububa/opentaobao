package alihouse

// AstoreSceneRespInfoDto 结构体
type AstoreSceneRespInfoDto struct {
	// 1
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 1
	OuterStoreId string `json:"outer_store_id,omitempty" xml:"outer_store_id,omitempty"`
	// 1
	OuterSellerId string `json:"outer_seller_id,omitempty" xml:"outer_seller_id,omitempty"`
	// 提示语
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// 1
	OuterType int64 `json:"outer_type,omitempty" xml:"outer_type,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
