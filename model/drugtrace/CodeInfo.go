package drugtrace

// CodeInfo 
type CodeInfo struct {
    // 码状态
    Status   string `json:"status,omitempty" xml:"status,omitempty"`
    // 追溯码
    Code   string `json:"code,omitempty" xml:"code,omitempty"`
    // 码等级
    CodeLevel   string `json:"code_level,omitempty" xml:"code_level,omitempty"`
    // 父码
    ParentCode   string `json:"parent_code,omitempty" xml:"parent_code,omitempty"`
    // 码包装等级
    CodePackLevel   string `json:"code_pack_level,omitempty" xml:"code_pack_level,omitempty"`
}
