package trade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoLifeReservationTradeConsumeNoticeAPIRequest
生服购后预约外部核销 API请求
taobao.life.reservation.trade.consume.notice

生服团购商品，购后预约。外部ISV进行核销 */
type TaobaoLifeReservationTradeConsumeNoticeAPIRequest struct {
	model.Params
	// 淘宝主单号
	_tradeNo string
	// 凭证ID
	_ticketId string
}

// New
