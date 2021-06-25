package gameact

import (
    "github.com/bububa/opentaobao/model"
)

/* 
抽奖 APIResponse
taobao.de.activity.luckydraw

用于激励平台对外提供抽奖功能，包括但不限于集分宝、红包、宝点、淘金币、淘彩票等
*/
type TaobaoDeActivityLuckydrawAPIResponse struct {
    model.CommonResponse
    Response *TaobaoDeActivityLuckydrawResponse `json:"taobao_de_activity_luckydraw_response,omitempty"`
}

type TaobaoDeActivityLuckydrawResponse struct {

    // 中奖奖品
    AwardVO   *AwardVO `json:"award_v_o,omitempty"`

    // 123
    IsLucky   bool `json:"is_lucky,omitempty"`

    // 数娱积分/金牌余额
    TotalAmount   int64 `json:"total_amount,omitempty"`

    // 可用抽奖次数
    AccessAmount   int64 `json:"access_amount,omitempty"`

}
