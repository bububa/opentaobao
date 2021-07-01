package trade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoLifeReservationItemOrderConfirmAPIRequest
生服购后预约单外部确认 API请求
taobao.life.reservation.item.order.confirm

生服团购下单预约后，用户改期/取消，外调ISV。ISV人工确认后调接口同意或驳回 */
type TaobaoLifeReservationItemOrderConfirmAPIRequest struct {
	model.Params
	// 淘宝主单号
	_tradeNo string
	// 凭证ID
	_ticketId string
	// 审核类型，PASS-通过；REJECT-驳回
	_optType string
}

// New
