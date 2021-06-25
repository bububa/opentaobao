package promotion

// AlibabaInteractLotterydrawDodrawResultDto 
type AlibabaInteractLotterydrawDodrawResultDto struct {

    // result
    LotteryDrawResult   *LotteryDrawResultDto `json:"lottery_draw_result,omitempty"`

    // code
    Code   int64 `json:"code,omitempty"`

    // success
    Success   bool `json:"success,omitempty"`

    // msg
    Msg   string `json:"msg,omitempty"`

}
