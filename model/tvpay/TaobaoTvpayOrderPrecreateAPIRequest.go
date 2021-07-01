package tvpay

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTvpayOrderPrecreateAPIRequest
tv支付预下单 API请求
taobao.tvpay.order.precreate

tv支付预下单 */
type TaobaoTvpayOrderPrecreateAPIRequest struct {
	model.Params
	// 设备id
	_deviceId string
	// 来源
	_from string
	// 订单详情
	_data string
	// 牌照方
	_license string
}

// New
