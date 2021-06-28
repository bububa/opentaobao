package film

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
淘票票抽奖活动查询API(渠道) APIResponse
taobao.film.lottery.rule.query

淘票票抽奖活动查询API，渠道维度查询
*/
type TaobaoFilmLotteryRuleQueryAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"film_lottery_rule_query_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *ResultListModel `json:"result,omitempty" xml:"