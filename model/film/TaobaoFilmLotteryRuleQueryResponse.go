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
    // Response *TaobaoFilmLotteryRuleQueryResponse `json:"film_lottery_rule_query_response,omitempty"` 
    TaobaoFilmLotteryRuleQueryResponse
}

/* model for simplify = false
type TaobaoFilmLotteryRuleQueryResponse struct {

    // result
    
    Result  *struct {
        ResultListModel  *ResultListModel `json:"result_list_model,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TaobaoFilmLotteryRuleQueryResponse struct {

    // result
    Result   *ResultListModel `json:"result,omitempty"`

}
