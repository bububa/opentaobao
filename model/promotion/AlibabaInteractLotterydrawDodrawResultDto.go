package promotion

// AlibabaInteractLotterydrawDodrawResultDTO 
type AlibabaInteractLotterydrawDodrawResultDTO struct {
    // result
    LotteryDrawResult   *LotteryDrawResultDTO `json:"lottery_draw_result,omitempty" xml:"lottery_draw_result,omitempty"`
    // code
    Code   int64 `json:"code,omitempty" xml:"code,omitempty"`
    // success
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    // msg
    Msg   string `json:"msg,omitempty" xml:"msg,omitempty"`
}
