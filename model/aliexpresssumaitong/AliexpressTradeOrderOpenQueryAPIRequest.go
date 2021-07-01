package aliexpresssumaitong

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AliexpressTradeOrderOpenQueryAPIRequest
Aliexpress开放平台订单查询 API请求
aliexpress.trade.order.open.query

Aliexpress开放平台订单信息查询 */
type AliexpressTradeOrderOpenQueryAPIRequest struct {
	model.Params
	// 买家用户id
	_buyerId int64
	// 订单号
	_orderIds []int64
	// 外部订单号
	_outIds []string
	// 小程序appId
	_openAppKey string
	// 业务编码
	_bizCode string
}

// New
