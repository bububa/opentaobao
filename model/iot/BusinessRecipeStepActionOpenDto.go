package iot

// BusinessRecipeStepActionOpenDto 
/* model for simplify = false
type BusinessRecipeStepActionOpenDto struct {

    // 指令名
    
    ActionName   string `json:"action_name,omitempty"`
    

    // 指令类型：0:整数型,1:枚举类型,2:浮点类型,3:布尔型,4:字符串,5:时间型,6:JSON对象
    
    ActionType   int64 `json:"action_type,omitempty"`
    

    // 指令值
    
    ActionValue   string `json:"action_value,omitempty"`
    

    // 指令顺序号
    
    Sequence   int64 `json:"sequence,omitempty"`
    

}
*/

// BusinessRecipeStepActionOpenDto 
type BusinessRecipeStepActionOpenDto struct {

    // 指令名
    ActionName   string `json:"action_name,omitempty"`

    // 指令类型：0:整数型,1:枚举类型,2:浮点类型,3:布尔型,4:字符串,5:时间型,6:JSON对象
    ActionType   int64 `json:"action_type,omitempty"`

    // 指令值
    ActionValue   string `json:"action_value,omitempty"`

    // 指令顺序号
    Sequence   int64 `json:"sequence,omitempty"`

}
