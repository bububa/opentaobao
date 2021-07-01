package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripMerchantGalaxyOrderQueryAPIRequest
星河-单个订单详细信息查询 API请求
alitrip.merchant.galaxy.order.query

为用户提供酒店订单的详细信息查询能力 */
type AlitripMerchantGalaxyOrderQueryAPIRequest struct {
	model.Params
	// 租户标识
	_tenantKey string
	// 用户登录标识
	_token string
	// 订单号
	_orderId string
}

// New
