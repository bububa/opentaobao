package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripMerchantGalaxyOrderCancelAPIRequest
星河-取消预订 API请求
alitrip.merchant.galaxy.order.cancel

雅高酒店用户使用该接口，取消酒店预订 */
type AlitripMerchantGalaxyOrderCancelAPIRequest struct {
	model.Params
	// 租户标识
	_tenantKey string
	// 用户登录标识
	_token string
	// 订单编号
	_orderId string
	// 退款原因
	_reason string
}

// New
