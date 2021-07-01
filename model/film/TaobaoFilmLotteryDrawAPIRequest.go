package film

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoFilmLotteryDrawAPIRequest
淘票票抽奖发放权益API API请求
taobao.film.lottery.draw

对外第三方合作渠道通过抽奖形式发码 */
type TaobaoFilmLotteryDrawAPIRequest struct {
	model.Params
	// 账号ID
	_accountNo string
	// 账号类型（TAOBAO\ALIPAY\PHONE\TAOBAO_NAME\OPEN_ID）
	_accountType string
	// 活动ID
	_lotteryMixId string
	// 平台类型
	_platform int64
	// 扩展参数
	_bizData string
}

// New
