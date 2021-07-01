package eticket

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoVmarketEticketFlowConsumeAPIRequest
无交易类凭证核销 API请求
taobao.vmarket.eticket.flow.consume

无交易类凭证核销 */
type TaobaoVmarketEticketFlowConsumeAPIRequest struct {
	model.Params
	// 业务单号
	_outerId string
	// 凭证码
	_code string
	// 淘宝业务提供的业务类型值，请联系相关业务运营取得
	_bizType int64
	// 核销操作人
	_operator string
}

// New
