package ottpay

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YoukuOttPayOrderDeleteorderAPIRequest
退订应用中心支付订单 API请求
youku.ott.pay.order.deleteorder

应用中心sdk连续包月退订接口 */
type YoukuOttPayOrderDeleteorderAPIRequest struct {
	model.Params
	// 下单账号， cp账号
	_buyer string
	// 商品id
	_productId string
	// 商品名称
	_productName string
	// cp订单号
	_orderNo string
	// 回调地址
	_callbackUrl string
	// 订单无关的其他参数,如埋点统计的utdid, mac地址等
	_extra string
	// 订单类型，1为连续包月类型,2为取消连续包月
	_orderType int64
	// 连续包月原始订单
	_originalOrderNo string
}

// New
