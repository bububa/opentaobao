package alitripmerchant

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
对外提供会员注册服务 APIResponse
alitrip.merchant.galaxy.member.provider.register

星河产品=对外提供注册雅高会员服务
*/
type AlitripMerchantGalaxyMemberProviderRegisterAPIResponse struct {
    model.CommonResponse
    AlitripMerchantGalaxyMemberProviderRegisterResponse
}

type AlitripMerchantGalaxyMemberProviderRegisterResponse struct {
    XMLName xml.Name `xml:"alitrip_merchant_galaxy_member_provider_register_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 默认描述
    
    Result   *Response `json:"result,omitempty" xml:"result,omitempty"`

    
}
