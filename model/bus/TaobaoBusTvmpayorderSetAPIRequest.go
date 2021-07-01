package bus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoBusTvmpayorderSetAPIRequest
自助机条形码被动支付 API请求
taobao.bus.tvmpayorder.set

汽车票线下自助机条形码支付 */
type TaobaoBusTvmpayorderSetAPIRequest struct {
	model.Params
	// 条形码认证码
	_alipayAuthCode string
	// 飞猪订单号
	_alitripOrderId string
}

// New
