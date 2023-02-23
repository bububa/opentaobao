package alihouse

// AstoreSceneInfoDto 结构体
type AstoreSceneInfoDto struct {
	// 外部主体id
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 外部门店id
	OuterStoreId string `json:"outer_store_id,omitempty" xml:"outer_store_id,omitempty"`
	// 外部商家id
	OuterSellerId string `json:"outer_seller_id,omitempty" xml:"outer_seller_id,omitempty"`
	// 外部主体类型
	OuterType int64 `json:"outer_type,omitempty" xml:"outer_type,omitempty"`
}
