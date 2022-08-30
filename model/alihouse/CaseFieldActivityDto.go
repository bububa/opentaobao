package alihouse

// CaseFieldActivityDto 结构体
type CaseFieldActivityDto struct {
	// 外部活动ID
	OuterActivityId string `json:"outer_activity_id,omitempty" xml:"outer_activity_id,omitempty"`
	// 活动名称
	ActivityName string `json:"activity_name,omitempty" xml:"activity_name,omitempty"`
	// 活动副标题
	ActivitySubName string `json:"activity_sub_name,omitempty" xml:"activity_sub_name,omitempty"`
	// 活动描述
	ActivityRemark string `json:"activity_remark,omitempty" xml:"activity_remark,omitempty"`
	// 活动规则
	ActivityRule string `json:"activity_rule,omitempty" xml:"activity_rule,omitempty"`
	// 活动展示开始时间
	ActivityBeginTime string `json:"activity_begin_time,omitempty" xml:"activity_begin_time,omitempty"`
	// 活动展示结束时间
	ActivityEndTime string `json:"activity_end_time,omitempty" xml:"activity_end_time,omitempty"`
	// 报名开始时间
	SignUpBeginTime string `json:"sign_up_begin_time,omitempty" xml:"sign_up_begin_time,omitempty"`
	// 报名结束时间
	SignUpEndTime string `json:"sign_up_end_time,omitempty" xml:"sign_up_end_time,omitempty"`
	// 活动举办开始时间
	HoldBeginTime string `json:"hold_begin_time,omitempty" xml:"hold_begin_time,omitempty"`
	// 活动举办结束时间
	HoldEndTime string `json:"hold_end_time,omitempty" xml:"hold_end_time,omitempty"`
	// 头图
	BannerUrls string `json:"banner_urls,omitempty" xml:"banner_urls,omitempty"`
	// 内容图
	ContentUrls string `json:"content_urls,omitempty" xml:"content_urls,omitempty"`
	// 联系人
	Contactor string `json:"contactor,omitempty" xml:"contactor,omitempty"`
	// 联系人电话
	ContactorPhone string `json:"contactor_phone,omitempty" xml:"contactor_phone,omitempty"`
	// 联系人邮箱
	ContactorEmail string `json:"contactor_email,omitempty" xml:"contactor_email,omitempty"`
	// 活动是否上架
	IsOnline int64 `json:"is_online,omitempty" xml:"is_online,omitempty"`
	// 是否在官盘透出活动
	IsOfficialShow int64 `json:"is_official_show,omitempty" xml:"is_official_show,omitempty"`
	// 是否签署免责声明
	IsSignDisclaimer int64 `json:"is_sign_disclaimer,omitempty" xml:"is_sign_disclaimer,omitempty"`
}
