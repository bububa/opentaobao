package wdk

// StairGroupDto 
/* model for simplify = false
type StairGroupDto struct {

    // 分组序号
    
    Number   int64 `json:"number,omitempty"`
    

    // 优惠门槛
    
    Condition  *struct {
        Condition  *Condition `json:"condition,omitempty"`
    } `json:"condition,omitempty"`
    

    // 优惠效果
    
    Action  *struct {
        Action  *Action `json:"action,omitempty"`
    } `json:"action,omitempty"`
    

}
*/

// StairGroupDto 
type StairGroupDto struct {

    // 分组序号
    Number   int64 `json:"number,omitempty"`

    // 优惠门槛
    Condition   *Condition `json:"condition,omitempty"`

    // 优惠效果
    Action   *Action `json:"action,omitempty"`

}
