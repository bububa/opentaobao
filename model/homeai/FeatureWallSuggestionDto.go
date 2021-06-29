package homeai

// FeatureWallSuggestionDto 
type FeatureWallSuggestionDto struct {

    // designid
    
    DesignId   string `json:"design_id,omitempty" xml:"design_id,omitempty"`
    

    // 墙位置
    
    FeatureWalls   []FeatureWallDto `json:"feature_walls,omitempty" xml:"feature_walls,omitempty"`
    

}
