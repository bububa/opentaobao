package alihouse

// BrokerMigrateDto 结构体
type BrokerMigrateDto struct {
	// 1
	OuterStoreId string `json:"outer_store_id,omitempty" xml:"outer_store_id,omitempty"`
	// 1
	NewOuterId string `json:"new_outer_id,omitempty" xml:"new_outer_id,omitempty"`
	// 1
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 1
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
	// 1
	IsRollBack int64 `json:"is_roll_back,omitempty" xml:"is_roll_back,omitempty"`
}
