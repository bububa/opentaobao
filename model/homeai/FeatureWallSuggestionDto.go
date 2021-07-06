package homeai

// FeatureWallSuggestionDto 结构体
type FeatureWallSuggestionDto struct {
	// 墙位置
	FeatureWalls []FeatureWallDto `json:"feature_walls,omitempty" xml:"feature_walls>feature_wall_dto,omitempty"`
	// designid
	DesignId string `json:"design_id,omitempty" xml:"design_id,omitempty"`
}
