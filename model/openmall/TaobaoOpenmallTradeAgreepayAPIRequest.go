package openmall

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOpenmallTradeAgreepayAPIRequest
openmall订单支付 API请求
taobao.openmall.trade.agreepay

openmall订单支付 */
type TaobaoOpenmallTradeAgreepayAPIRequest struct {
	model.Params
	// 媒体渠道，代表分销者的身份，签约支付宝代扣的渠道商淘宝账号nick
	_distributor string
	// 淘宝交易单号
	_tid int64
}

// New
