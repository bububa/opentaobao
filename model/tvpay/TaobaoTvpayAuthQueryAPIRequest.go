package tvpay

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTvpayAuthQueryAPIRequest
tv支付授权查询 API请求
taobao.tvpay.auth.query

查询该用户在指定设备上是否有支付授权 */
type TaobaoTvpayAuthQueryAPIRequest struct {
	model.Params
	// 设备号
	_deviceId string
	// 来源
	_from string
	// 业务订单号
	_bizOrderId string
	// 是否淘系
	_isTao bool
	// 支付宝订单号
	_orderNo string
	// 外部订单号
	_outOrderNo string
}

// New
