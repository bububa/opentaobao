package alihouse

// StoreTalentBindingResponseDto 结构体
type StoreTalentBindingResponseDto struct {
	// 外部id
	TargetId string `json:"target_id,omitempty" xml:"target_id,omitempty"`
	// 达人号
	TalentId string `json:"talent_id,omitempty" xml:"talent_id,omitempty"`
	// 错误原因
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// 账号类型
	AccountType int64 `json:"account_type,omitempty" xml:"account_type,omitempty"`
	// 是否绑定成功
	IsSuccess int64 `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
