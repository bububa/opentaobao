package drugtrace

// CodeInfoListDTO 
type CodeInfoListDTO struct {
    // 制剂规格
    PrepnSpec   string `json:"prepn_spec,omitempty" xml:"prepn_spec,omitempty"`
    // 最小制剂数量
    PrepnAmount   string `json:"prepn_amount,omitempty" xml:"prepn_amount,omitempty"`
    // 最小包装数量
    PkgAmount   string `json:"pkg_amount,omitempty" xml:"pkg_amount,omitempty"`
    // 监管码级别
    CodeLevel   string `json:"code_level,omitempty" xml:"code_level,omitempty"`
    // 监管码
    Code   string `json:"code,omitempty" xml:"code,omitempty"`
}
