package alitripmerchant

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
星河-支付参数查询接口 APIResponse
alitrip.merchant.galaxy.payment.param.query

获取微信小程序的支付参数，供微信小程序调起支付请求。参数都为校验字段，不涉及用户隐私数据
*/
type AlitripMerchantGalaxyPaymentParamQueryAPIResponse struct {
    model.CommonResponse
    AlitripMerchantGalaxyPaymentParamQueryResponse
}

type AlitripMerchantGalaxyPaymentParamQueryResponse struct {
    XMLName xml.Name `xml:"alitrip_merchant_galaxy_payment_param_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 默认描述
    
    Result   *Response `json:"result,omitempty" xml:"result,omitempty"`

    
}
