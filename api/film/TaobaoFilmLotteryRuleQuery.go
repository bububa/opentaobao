package film

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/film"
)

/* 
淘票票抽奖活动查询API(渠道) 
taobao.film.lottery.rule.query

淘票票抽奖活动查询API，渠道维度查询
*/
func TaobaoFilmLotteryRuleQuery(clt *core.SDKClient, req *film.TaobaoFilmLotteryRuleQueryRequest, session string) (*film.TaobaoFilmLotteryRuleQueryAPIResponse, error) {
    var resp film.TaobaoFilmLotteryRuleQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
