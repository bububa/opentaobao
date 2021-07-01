package film

// TaobaoFilmLotteryDrawModel 结构体
type TaobaoFilmLotteryDrawModel struct {
	// 奖品模型
	Rewards []DrawReward `json:"rewards,omitempty" xml:"rewards>draw_reward,omitempty"`
	// 使用地址
	UseUrl string `json:"use_url,omitempty" xml:"use_url,omitempty"`
}
