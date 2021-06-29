package legalsuit

// CourtEvidenceModel 
type CourtEvidenceModel struct {

    // 质证意见
    
    Opinion   string `json:"opinion,omitempty" xml:"opinion,omitempty"`
    

    // 是否认可
    
    IsAgreed   string `json:"is_agreed,omitempty" xml:"is_agreed,omitempty"`
    

    // 质证人
    
    Evidence   string `json:"evidence,omitempty" xml:"evidence,omitempty"`
    

    // 证据目的
    
    EvidenceAim   string `json:"evidence_aim,omitempty" xml:"evidence_aim,omitempty"`
    

    // 证据名称
    
    EvidenceName   string `json:"evidence_name,omitempty" xml:"evidence_name,omitempty"`
    

    // 姓名
    
    Name   string `json:"name,omitempty" xml:"name,omitempty"`
    

    // 质证证据类型(原告证据还是被告证据)
    
    Type   string `json:"type,omitempty" xml:"type,omitempty"`
    

}
