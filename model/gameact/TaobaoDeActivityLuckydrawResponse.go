package gameact

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
抽奖 APIResponse
taobao.de.activity.luckydraw

用于激励平台对外提供抽奖功能，包括但不限于集分宝、红包、宝点、淘金币、淘彩票等
*/
type TaobaoDeActivityLuckydrawAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"de_activity_luckydraw_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 中奖奖品
    
    AwardVO   *AwardVO `json:"award_v_o,omitempty" xml:"