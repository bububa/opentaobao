package alitripmerchant

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
星河-校验token APIResponse
alitrip.merchant.galaxy.member.token

校验或者刷新token
*/
type AlitripMerchantGalaxyMemberTokenAPIResponse struct {
    model.CommonResponse
    AlitripMerchantGalaxyMemberTokenResponse
}

type AlitripMerchantGalaxyMemberTokenResponse struct {
    XMLName xml.Name `xml:"alitrip_merchant_galaxy_member_token_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 默认描述
    
    Result   *Response `json:"result,omitempty" xml:"result,omitempty"`

    
}
