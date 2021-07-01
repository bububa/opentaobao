package nrt

// StallSigningRespDto 结构体
type StallSigningRespDto struct {
	// 申请单id
	OrderId string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 摊位/门店id
	StoreId int64 `json:"store_id,omitempty" xml:"store_id,omitempty"`
}
