package film

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoFilmLotteryRuleQueryAPIRequest
淘票票抽奖活动查询API(渠道) API请求
taobao.film.lottery.rule.query

淘票票抽奖活动查询API，渠道维度查询 */
type TaobaoFilmLotteryRuleQueryAPIRequest struct {
	model.Params
	// 账号类型（TAOBAO\ALIPAY\PHONE）
	_accountType string
	// 渠道来源
	_channel string
	// 账号ID
	_accountNo string
}

// New
