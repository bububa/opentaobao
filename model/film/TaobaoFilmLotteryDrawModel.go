package film

// TaobaoFilmLotteryDrawModel 
/* model for simplify = false
type TaobaoFilmLotteryDrawModel struct {

    // 奖品模型
    
    Rewards  struct {
        DrawReward  []DrawReward `json:"draw_reward,omitempty"`
    } `json:"rewards,omitempty"`
    

    // 使用地址
    
    UseUrl   string `json:"use_url,omitempty"`
    

}
*/

// TaobaoFilmLotteryDrawModel 
type TaobaoFilmLotteryDrawModel struct {

    // 奖品模型
    Rewards   []DrawReward `json:"rewards,omitempty"`

    // 使用地址
    UseUrl   string `json:"use_url,omitempty"`

}
