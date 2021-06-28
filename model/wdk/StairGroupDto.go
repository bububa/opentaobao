package wdk

// StairGroupDto 
type StairGroupDto struct {

    // 分组序号
    
    Number   int64 `json:"number,omitempty" xml:"number,omitempty"`
    

    // 优惠门槛
    
    Condition   *Condition `json:"condition,omitempty" xml:"condition,omitempty"`
    

    // 优惠效果
    
    Action   *Action `json:"action,omitempty" xml:"action,omitempty"`
    

}
