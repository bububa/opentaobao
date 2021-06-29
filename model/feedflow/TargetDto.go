package feedflow

// TargetDto 
type TargetDto struct {

    // 定向id
    
    TargetId   int64 `json:"target_id,omitempty" xml:"target_id,omitempty"`
    

    // 定向名称
    
    TargetName   string `json:"target_name,omitempty" xml:"target_name,omitempty"`
    

    // 定向描述
    
    TargetDesc   string `json:"target_desc,omitempty" xml:"target_desc,omitempty"`
    

    // 定向类型
    
    TargetType   string `json:"target_type,omitempty" xml:"target_type,omitempty"`
    

}
