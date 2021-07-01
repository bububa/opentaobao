package ottpay

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YoukuOttPayOrderCreateorderAPIRequest
创建订单 API请求
youku.ott.pay.order.createorder

ottpay创建订单 */
type YoukuOttPayOrderCreateorderAPIRequest struct {
	model.Params
	// 下单账号， cp账号
	_buyer string
	// 商品id
	_productId string
	// 商品名称
	_productName string
	// cp订单号
	_orderNo string
	// 价格， 单位：分
	_price string
	// 回调接口
	_callbackUrl string
	// 订单无关的其他参数,如埋点统计的utdid, mac地址等
	_extra string
	// 订单类型，1为连续包月类型
	_orderType int64
	// 连续包月实际参数
	_realPrice string
}

// New
