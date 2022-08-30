package alihealth2

// IMRule 结构体
type IMRule struct {
	// INTEGRITY("对话完整性"),     ATTITUDE("态度"),     CONTENT_SECURITY("内容安全")
	Type string `json:"type,omitempty" xml:"type,omitempty"`
	// 只有ATTITUDE类型才有值，识别的结果，如："语气问题", "辱骂骚扰"，"自我介绍", "感谢用语", "祝福用语",      * "礼貌用语", "患者安抚", "随访追问", "邀评关注", "平台推荐"等
	Label string `json:"label,omitempty" xml:"label,omitempty"`
	// 1: 违规， 2: 疑似， -1: 该类型下，不对结果进行评判
	Res int64 `json:"res,omitempty" xml:"res,omitempty"`
}
