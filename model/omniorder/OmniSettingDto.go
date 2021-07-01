package omniorder

// OmniSettingDto 结构体
type OmniSettingDto struct {
	// 接单系统，填 0 代表店掌柜，填 1 代表 POS
	AcceptedSystem int64 `json:"accepted_system,omitempty" xml:"accepted_system,omitempty"`
	// 分单系统，填 astrolabe 代表阿里分单，填 RDS的 appkey 代表自行分单
	AllocatedSystem string `json:"allocated_system,omitempty" xml:"allocated_system,omitempty"`
}
