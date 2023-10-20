package alitripmerchant

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripmerchantgalaxypaymentparamqueryAPIResponse 星河-支付参数查询接口 API返回值
// alitrip.merchant.galaxy.payment.param.query
//
// 获取微信小程序的支付参数，供微信小程序调起支付请求。参数都为校验字段，不涉及用户隐私数据
type AlitripmerchantgalaxypaymentparamqueryAPIResponse struct {
	model.CommonResponse
	AlitripmerchantgalaxypaymentparamqueryAPIResponseModel
}

// AlitripmerchantgalaxypaymentparamqueryAPIResponseModel is 星河-支付参数查询接口 成功返回结果
type AlitripmerchantgalaxypaymentparamqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_payment_param_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 默认描述
	Result *AlitripmerchantgalaxypaymentparamqueryResponse `json:"result,omitempty" xml:"result,omitempty"`
}
