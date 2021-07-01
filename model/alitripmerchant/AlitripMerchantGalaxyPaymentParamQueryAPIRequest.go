package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripMerchantGalaxyPaymentParamQueryAPIRequest
星河-支付参数查询接口 API请求
alitrip.merchant.galaxy.payment.param.query

获取微信小程序的支付参数，供微信小程序调起支付请求。参数都为校验字段，不涉及用户隐私数据 */
type AlitripMerchantGalaxyPaymentParamQueryAPIRequest struct {
	model.Params
	// 租户身份信息
	_tenantKey string
	// 用户校验token
	_token string
	// 订单编号
	_orderId string
}

// New
