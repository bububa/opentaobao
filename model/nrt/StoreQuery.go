package nrt

// StoreQuery 结构体
type StoreQuery struct {
	// 卖场Id或者同城id
	StoreIds []int64 `json:"store_ids,omitempty" xml:"store_ids>int64,omitempty"`
	// 类型
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
}
