package film

import (
    "github.com/bububa/opentaobao/model"
)

/* 
淘票票抽奖活动查询API(渠道) APIResponse
taobao.film.lottery.rule.query

淘票票抽奖活动查询API，渠道维度查询
*/
type TaobaoFilmLotteryRuleQueryAPIResponse struct {
    model.CommonResponse
    Response *TaobaoFilmLotteryRuleQueryResponse `json:"taobao_film_lottery_rule_query_response,omitempty"`
}

type TaobaoFilmLotteryRuleQueryResponse struct {

    // result
    Result   *ResultListModel `json:"result,omitempty"`

}
