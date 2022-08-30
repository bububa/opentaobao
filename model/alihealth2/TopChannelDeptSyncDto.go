package alihealth2

// TopChannelDeptSyncDto 结构体
type TopChannelDeptSyncDto struct {
	// 医院ID+科室ID+状态
	BizContent string `json:"biz_content,omitempty" xml:"biz_content,omitempty"`
}
