package homeai

// FeatureWallSuggestionDTO 
type FeatureWallSuggestionDTO struct {
    // designid
    DesignId   string `json:"design_id,omitempty" xml:"design_id,omitempty"`
    // 墙位置
    FeatureWalls   []FeatureWallDTO `json:"feature_walls,omitempty" xml:"feature_walls>feature_wall_dto,omitempty"`
}
