package alitripmerchant

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
星河-获取登录用户的信息 APIResponse
alitrip.merchant.galaxy.member.query

获取登录用户的信息
*/
type AlitripMerchantGalaxyMemberQueryAPIResponse struct {
    model.CommonResponse
    AlitripMerchantGalaxyMemberQueryResponse
}

type AlitripMerchantGalaxyMemberQueryResponse struct {
    XMLName xml.Name `xml:"alitrip_merchant_galaxy_member_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 默认描述
    
    Result   *Response `json:"result,omitempty" xml:"result,omitempty"`

    
}
