package alitripmerchant

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
星河-微信小程序会员注册 APIResponse
alitrip.merchant.galaxy.member.register

星河产品=微信小程序注册雅高会员服务
*/
type AlitripMerchantGalaxyMemberRegisterAPIResponse struct {
    model.CommonResponse
    AlitripMerchantGalaxyMemberRegisterResponse
}

type AlitripMerchantGalaxyMemberRegisterResponse struct {
    XMLName xml.Name `xml:"alitrip_merchant_galaxy_member_register_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 默认描述
    
    Result   *Response `json:"result,omitempty" xml:"result,omitempty"`

    
}
