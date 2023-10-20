package alihealth2

// ImriskCheckCommand 结构体
type ImriskCheckCommand struct {
	// 会话内容，按照时间排序
	Conversations []Conversation `json:"conversations,omitempty" xml:"conversations>conversation,omitempty"`
	// 患者id
	PatientId string `json:"patient_id,omitempty" xml:"patient_id,omitempty"`
	// 医生id
	DoctorId string `json:"doctor_id,omitempty" xml:"doctor_id,omitempty"`
	// 场景
	SceneName string `json:"scene_name,omitempty" xml:"scene_name,omitempty"`
	// 租户
	TenantCode string `json:"tenant_code,omitempty" xml:"tenant_code,omitempty"`
	// 会话id
	SessionId string `json:"session_id,omitempty" xml:"session_id,omitempty"`
	// 会话开始时间，YYYY-MM-DD HH:mm:ss格式
	BizTime string `json:"biz_time,omitempty" xml:"biz_time,omitempty"`
}
