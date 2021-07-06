package tanx

// QualificationQuery 结构体
type QualificationQuery struct {
	// tanx系统广告主表userId，查询时和user_names选其一
	UserIds []int64 `json:"user_ids,omitempty" xml:"user_ids>int64,omitempty"`
	// 广告主名称，淘系用户为旺旺名称，非淘系客户为营业执照上的名称。查询时和user_ids请只选其一
	UserNames []string `json:"user_names,omitempty" xml:"user_names>string,omitempty"`
	// 查询时包含的资质元素id列表
	ElementIds []int64 `json:"element_ids,omitempty" xml:"element_ids>int64,omitempty"`
	// 资质元素id
	Ids []int64 `json:"ids,omitempty" xml:"ids>int64,omitempty"`
	// 生效时间左区间
	StartTimeBegin string `json:"start_time_begin,omitempty" xml:"start_time_begin,omitempty"`
	// 排序字段
	OrderBy string `json:"order_by,omitempty" xml:"order_by,omitempty"`
	// 失效时间左区间
	EndTimeEnd string `json:"end_time_end,omitempty" xml:"end_time_end,omitempty"`
	// 生效时间右区间
	StartTimeEnd string `json:"start_time_end,omitempty" xml:"start_time_end,omitempty"`
	// 创建时间右区间
	CreateTimeEnd string `json:"create_time_end,omitempty" xml:"create_time_end,omitempty"`
	// 审核时间左区间
	AuditTimeBegin string `json:"audit_time_begin,omitempty" xml:"audit_time_begin,omitempty"`
	// 审核时间右区间
	AuditTimeEnd string `json:"audit_time_end,omitempty" xml:"audit_time_end,omitempty"`
	// 失效时间右区间
	EndTimeBegin string `json:"end_time_begin,omitempty" xml:"end_time_begin,omitempty"`
	// 创建时间左区间
	CreateTimeBegin string `json:"create_time_begin,omitempty" xml:"create_time_begin,omitempty"`
	// 资质生效状态(该状态值是根据不能在新增资质时设置)-1=已过期，0=待生效，1=生效中，2=即将过期
	EffectiveStatus int64 `json:"effective_status,omitempty" xml:"effective_status,omitempty"`
	// 客户类型 1-淘系，2-非淘系dsp公司，3-非淘系dsp个人
	UserType int64 `json:"user_type,omitempty" xml:"user_type,omitempty"`
	// 审核状态 -1=拒绝，0=待审核，1=通过
	AuditStatus int64 `json:"audit_status,omitempty" xml:"audit_status,omitempty"`
	// 正序还是倒序 1是正,0是倒
	Asc int64 `json:"asc,omitempty" xml:"asc,omitempty"`
}
