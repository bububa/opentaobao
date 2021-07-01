package openmall

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOpenmallTradeConfirmAPIRequest
确认收货 API请求
taobao.openmall.trade.confirm

确认订单收货 */
type TaobaoOpenmallTradeConfirmAPIRequest struct {
	model.Params
	// 分销者信息
	_distributor string
	// 淘宝订单号
	_tid int64
}

// New
