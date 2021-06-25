package alsc

// CodeBizAttributeDto 
type CodeBizAttributeDto struct {

    // 码值
    CodeValue   string `json:"code_value,omitempty"`

    // 码值绑定业务主体类型
    SubjectType   string `json:"subject_type,omitempty"`

    // 码值绑定的业务主体ID
    SubjectId   string `json:"subject_id,omitempty"`

}
