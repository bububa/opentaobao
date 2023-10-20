package alihealth2

// CommonRequest4top 结构体
type CommonRequest4top struct {
	// hosNo:渠道医院ID,deptNo:渠道科室ID,deptName:科室名称,status:状态
	BizContent string `json:"biz_content,omitempty" xml:"biz_content,omitempty"`
}
