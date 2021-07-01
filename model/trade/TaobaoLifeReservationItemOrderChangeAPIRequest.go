package trade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoLifeReservationItemOrderChangeAPIRequest
生服购后预约单外部发起变更 API请求
taobao.life.reservation.item.order.change

生服购后预约单外部发起变更，例如改期、取消。目前体检场景，用户会直接联系ISV改期/取消，因此开放给ISV这块的能力 */
type TaobaoLifeReservationItemOrderChangeAPIRequest struct {
	model.Params
	// 淘宝主单号
	_tradeNo string
	// 凭证ID
	_ticketId string
	// 改期：MODIFY   取消：CANCEL
	_action string
	// 改期必填，格式：yyyy-MM-dd HH:mm。时分固定00:00
	_reserveStartTime string
	// 改期必填，格式：yyyy-MM-dd HH:mm。时分固定23:59
	_reserveEndTime string
}

// New
