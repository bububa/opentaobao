package film

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoFilmLotteryRuleQueryAPIResponse
淘票票抽奖活动查询API(渠道) API返回值
taobao.film.lottery.rule.query

淘票票抽奖活动查询API，渠道维度查询 */
type TaobaoFilmLotteryRuleQueryAPIResponse struct {
	model.CommonResponse
	TaobaoFilmLotteryRuleQueryAPIResponseModel
}

// TaobaoFilmLotteryRuleQueryAPIResponseModel is 淘票票抽奖活动查询API(渠道) 成功返回结果
type TaobaoFilmLotteryRuleQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"film_lottery_rule_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *ResultListModel `json:"result,omitempty" xml:"result,omitempty"`
}
