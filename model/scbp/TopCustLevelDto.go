package scbp

// TopCustLevelDto 
type TopCustLevelDto struct {
    // levelScore
    LevelScore   int64 `json:"level_score,omitempty" xml:"level_score,omitempty"`
    // growthLevel
    GrowthLevel   string `json:"growth_level,omitempty" xml:"growth_level,omitempty"`
}
