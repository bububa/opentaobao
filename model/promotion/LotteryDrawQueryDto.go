package promotion

// LotteryDrawQueryDto 结构体
type LotteryDrawQueryDto struct {
	// 抽奖参数
	Wua string `json:"wua,omitempty" xml:"wua,omitempty"`
	// 关联ID
	RelationId string `json:"relation_id,omitempty" xml:"relation_id,omitempty"`
	// 抽奖CODE
	RaffleCode string `json:"raffle_code,omitempty" xml:"raffle_code,omitempty"`
	// 前台appKey
	RaffleIdentity string `json:"raffle_identity,omitempty" xml:"raffle_identity,omitempty"`
	// 方案ID
	SchemaId int64 `json:"schema_id,omitempty" xml:"schema_id,omitempty"`
	// 抽奖参数
	UmidToken string `json:"umid_token,omitempty" xml:"umid_token,omitempty"`
	// 抽奖参数
	Ua string `json:"ua,omitempty" xml:"ua,omitempty"`
	// 用户终端信息
	UserAgent string `json:"user_agent,omitempty" xml:"user_agent,omitempty"`
	// 用户ip
	RemoteIp string `json:"remote_ip,omitempty" xml:"remote_ip,omitempty"`
	// 买家混淆Nick
	BuyerMixNick string `json:"buyer_mix_nick,omitempty" xml:"buyer_mix_nick,omitempty"`
}
