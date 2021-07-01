package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripMerchantGalaxyOrderBookAPIRequest
星河-订单预订接口 API请求
alitrip.merchant.galaxy.order.book

为星河酒店解决方案的C端用户提供酒店预订能力 */
type AlitripMerchantGalaxyOrderBookAPIRequest struct {
	model.Params
	// 租户身份信息
	_tenantKey string
	// 用户登录token
	_token string
	// 预订参数
	_orderParam *CreateOrderParam
	// 订单编号
	_orderCode string
	// 广告追踪信息
	_sourceQuery string
}

// New
