package promotion

// LotteryDrawQueryDto 
type LotteryDrawQueryDto struct {

    // 抽奖参数
    Wua   string `json:"wua,omitempty"`

    // 关联ID
    RelationId   string `json:"relation_id,omitempty"`

    // 抽奖CODE
    RaffleCode   string `json:"raffle_code,omitempty"`

    // 前台appKey
    RaffleIdentity   string `json:"raffle_identity,omitempty"`

    // 方案ID
    SchemaId   int64 `json:"schema_id,omitempty"`

    // 抽奖参数
    UmidToken   string `json:"umid_token,omitempty"`

    // 抽奖参数
    Ua   string `json:"ua,omitempty"`

    // 用户终端信息
    UserAgent   string `json:"user_agent,omitempty"`

    // 用户ip
    RemoteIp   string `json:"remote_ip,omitempty"`

    // 买家混淆Nick
    BuyerMixNick   string `json:"buyer_mix_nick,omitempty"`

}
