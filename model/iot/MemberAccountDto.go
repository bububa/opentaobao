package iot

// MemberAccountDTO 
type MemberAccountDTO struct {
    // 等级名称
    GradeName   string `json:"grade_name,omitempty" xml:"grade_name,omitempty"`
    // 等级编号
    Grade   int64 `json:"grade,omitempty" xml:"grade,omitempty"`
    // gmtCreate
    GmtCreate   string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
    // bindStatus 1：绑卡（已经是线下会员线上未绑定，或者解绑后再绑定），2：注册
    BindStatus   int64 `json:"bind_status,omitempty" xml:"bind_status,omitempty"`
}
