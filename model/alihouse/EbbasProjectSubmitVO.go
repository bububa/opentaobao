package alihouse

// EbbasProjectSubmitVO 
type EbbasProjectSubmitVO struct {
    // 楼盘code
    ProjectCode   string `json:"project_code,omitempty" xml:"project_code,omitempty"`
    // 楼盘e码
    ECode   string `json:"e_code,omitempty" xml:"e_code,omitempty"`
    // 外部id
    OuterId   string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
}
