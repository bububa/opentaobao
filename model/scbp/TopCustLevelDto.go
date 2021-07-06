package scbp

// TopCustLevelDto 结构体
type TopCustLevelDto struct {
	// growthLevel
	GrowthLevel string `json:"growth_level,omitempty" xml:"growth_level,omitempty"`
	// levelScore
	LevelScore int64 `json:"level_score,omitempty" xml:"level_score,omitempty"`
}
