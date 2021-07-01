package ottpay

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YoukuOttPayOrderAuthpayAPIRequest
委托代扣服务 API请求
youku.ott.pay.order.authpay

应用中心sdk连续包月委托代扣服务 */
type YoukuOttPayOrderAuthpayAPIRequest struct {
	model.Params
	// cp用户名
	_buyer string
	// 连续包月原始cp订单号
	_originalOrderNo string
	// 委托扣款cp订单号
	_orderNo string
	// 已配置开通连续包月的产品id
	_productId string
	// 回调
	_callbackUrl string
}

// New
