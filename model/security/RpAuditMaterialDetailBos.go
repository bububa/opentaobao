package security

// RpAuditMaterialDetailBos 
type RpAuditMaterialDetailBos struct {

    // 给用户的建议
    
    Suggestion   string `json:"suggestion,omitempty" xml:"suggestion,omitempty"`
    

    // 外化给用户的文案
    
    Display   string `json:"display,omitempty" xml:"display,omitempty"`
    

    // 是否安全问题
    
    Security   bool `json:"security,omitempty" xml:"security,omitempty"`
    

    // intercept
    
    Intercept   bool `json:"intercept,omitempty" xml:"intercept,omitempty"`
    

    // 材料结论文本描述
    
    Text   string `json:"text,omitempty" xml:"text,omitempty"`
    

    // 材料结论编码
    
    Code   string `json:"code,omitempty" xml:"code,omitempty"`
    

    // 材料结论分类
    
    Type   string `json:"type,omitempty" xml:"type,omitempty"`
    

    // 材料类别
    
    MaterialType   string `json:"material_type,omitempty" xml:"material_type,omitempty"`
    

}
