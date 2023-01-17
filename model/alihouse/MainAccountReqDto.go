package alihouse

// MainAccountReqDto 结构体
type MainAccountReqDto struct {
	// 账号主体外部id
	OuterTargetId string `json:"outer_target_id,omitempty" xml:"outer_target_id,omitempty"`
	// 主账号昵称
	MainUserNick string `json:"main_user_nick,omitempty" xml:"main_user_nick,omitempty"`
	// 是否测试 0-非测试 1-测试
	IsTest int64 `json:"is_test,omitempty" xml:"is_test,omitempty"`
	// 账号类型 1-公司 2-城市品牌公司 3-品牌 4-门店 5-签约公司
	AccountType int64 `json:"account_type,omitempty" xml:"account_type,omitempty"`
	// etc版本号
	EtcVersion int64 `json:"etc_version,omitempty" xml:"etc_version,omitempty"`
}
