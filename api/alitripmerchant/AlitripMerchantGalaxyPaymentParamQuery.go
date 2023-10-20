package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyPaymentParamQuery 星河-支付参数查询接口
// alitrip.merchant.galaxy.payment.param.query
//
// 获取微信小程序的支付参数，供微信小程序调起支付请求。参数都为校验字段，不涉及用户隐私数据
func AlitripMerchantGalaxyPaymentParamQuery(clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyPaymentParamQueryAPIRequest, resp *alitripmerchant.AlitripMerchantGalaxyPaymentParamQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
