package campus

// SpaceIdsQuery 结构体
type SpaceIdsQuery struct {
	// building/floor
	SpaceType string `json:"space_type,omitempty" xml:"space_type,omitempty"`
	// ids
	Ids []int64 `json:"ids,omitempty" xml:"ids>int64,omitempty"`
}
