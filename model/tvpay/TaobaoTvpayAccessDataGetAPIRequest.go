package tvpay

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTvpayAccessDataGetAPIRequest
tv支付 API请求
taobao.tvpay.access.data.get

在匿名用户支付后尝试为其登陆绑定的淘宝账号 */
type TaobaoTvpayAccessDataGetAPIRequest struct {
	model.Params
	// 设备id
	_deviceId string
	// 来源
	_from string
	// 订单id
	_outOrderNo string
	// 账号客户端版本
	_accountClientVersion string
}

// New
