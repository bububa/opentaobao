package tanx

// QualificationDto 结构体
type QualificationDto struct {
	// 通过的url
	Urls []string `json:"urls,omitempty" xml:"urls>string,omitempty"`
	// 通过的行业
	Specialindustrys []string `json:"specialindustrys,omitempty" xml:"specialindustrys>string,omitempty"`
	// 资质内容id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 资质名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 用户id
	UserId int64 `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// 旺旺或者dsp广告主名称
	UserName string `json:"user_name,omitempty" xml:"user_name,omitempty"`
	// 广告主类别（0-淘宝，1-天猫，2-dsp公司，3-dsp个人）
	UserType int64 `json:"user_type,omitempty" xml:"user_type,omitempty"`
	// 资质模板元素id
	ElementId int64 `json:"element_id,omitempty" xml:"element_id,omitempty"`
	// 资质内容列表
	ContentList []string `json:"content_list,omitempty" xml:"content_list>string,omitempty"`
	// -1=已过期，0=待生效，1=生效中，2=即将过期
	EffectiveStatus int64 `json:"effective_status,omitempty" xml:"effective_status,omitempty"`
	// 资质审核状态 -1=拒绝，0=待审核，1=通过
	AuditStatus int64 `json:"audit_status,omitempty" xml:"audit_status,omitempty"`
	// dspid,淘系内部产品也统一成dsp
	DspId int64 `json:"dsp_id,omitempty" xml:"dsp_id,omitempty"`
	// 资质生效时间
	StartTime string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 资质失效时间
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 资质创建时间
	CreateTime string `json:"create_time,omitempty" xml:"create_time,omitempty"`
	// 资质修改时间
	UpdateTime string `json:"update_time,omitempty" xml:"update_time,omitempty"`
	// 资质审核时间
	AuditTime string `json:"audit_time,omitempty" xml:"audit_time,omitempty"`
	// 拒绝原因
	Reason string `json:"reason,omitempty" xml:"reason,omitempty"`
	// 用户附加内容
	Supplement string `json:"supplement,omitempty" xml:"supplement,omitempty"`
	// 排查时间
	CheckTime string `json:"check_time,omitempty" xml:"check_time,omitempty"`
}
