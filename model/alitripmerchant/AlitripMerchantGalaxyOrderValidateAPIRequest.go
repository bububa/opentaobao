package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripMerchantGalaxyOrderValidateAPIRequest
星河-订单试单接口 API请求
alitrip.merchant.galaxy.order.validate

根据用户选择酒店房型、入住人数、预订时间参数，获取是否可预订及价格变化信息 */
type AlitripMerchantGalaxyOrderValidateAPIRequest struct {
	model.Params
	// 租户身份信息
	_tenantKey string
	// 试单参数
	_validateOrderParam *ValidateOrderParam
	// 用户标识
	_token string
}

// New
