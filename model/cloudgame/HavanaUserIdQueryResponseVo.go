package cloudgame

// HavanaUserIdQueryResponseVo 结构体
type HavanaUserIdQueryResponseVo struct {
	// 账号id
	AccountId string `json:"account_id,omitempty" xml:"account_id,omitempty"`
	// 缓存时间
	Ttl int64 `json:"ttl,omitempty" xml:"ttl,omitempty"`
	// 账号状态
	SessionState int64 `json:"session_state,omitempty" xml:"session_state,omitempty"`
	// 账号域
	AccountDomain int64 `json:"account_domain,omitempty" xml:"account_domain,omitempty"`
}
