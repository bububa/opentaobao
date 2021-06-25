package film

// TaobaoFilmLotteryDrawModel 
type TaobaoFilmLotteryDrawModel struct {

    // 奖品模型
    Rewards   []DrawReward `json:"rewards,omitempty"`

    // 使用地址
    UseUrl   string `json:"use_url,omitempty"`

}
