package admarket

// Target 
type Target struct {

    // 广告目标类型
    
    TargetType   string `json:"target_type,omitempty" xml:"target_type,omitempty"`
    

    // 广告目标值
    
    TargetValue   string `json:"target_value,omitempty" xml:"target_value,omitempty"`
    

}
