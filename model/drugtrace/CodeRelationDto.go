package drugtrace

// CodeRelationDTO 
type CodeRelationDTO struct {

    // 追溯码；查询的码
    
    Code   string `json:"code,omitempty" xml:"code,omitempty"`
    

    // 父码
    
    ParentCode   string `json:"parent_code,omitempty" xml:"parent_code,omitempty"`
    

    // 码级别
    
    CodeLevel   string `json:"code_level,omitempty" xml:"code_level,omitempty"`
    

    // 包装级别
    
    CodePackLevel   string `json:"code_pack_level,omitempty" xml:"code_pack_level,omitempty"`
    

    // 装箱数量;小盒码，返回1；中包码，返回实际小盒数量；大箱码，返回实际小盒数量
    
    BoxAmount   int64 `json:"box_amount,omitempty" xml:"box_amount,omitempty"`
    

    // 大箱或中包状态；若扫描的是小盒码，直接返回正常； 0-正常；1-拼箱；2-零箱；3-即拼箱又零箱
    
    BoxStatus   int64 `json:"box_status,omitempty" xml:"box_status,omitempty"`
    

    // 码状态
    
    Status   string `json:"status,omitempty" xml:"status,omitempty"`
    

}
