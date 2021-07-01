package tvpay

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTvpayOrderQueryAPIRequest
tv支付查询订单状态 API请求
taobao.tvpay.order.query

tv支付查询订单状态 */
type TaobaoTvpayOrderQueryAPIRequest struct {
	model.Params
	// 设备id
	_deviceId string
	// 来源
	_from string
	// 业务订单号
	_bizOrderId string
	// 是否淘系
	_isTao bool
	// 支付宝订单号
	_orderNo string
	// 订单类型
	_orderType string
	// 外部订单号
	_outOrderNo string
	// 牌照方
	_license string
}

// New
