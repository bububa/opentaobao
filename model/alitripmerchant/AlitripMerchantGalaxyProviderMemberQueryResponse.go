package alitripmerchant

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
提供会员查询接口 APIResponse
alitrip.merchant.galaxy.provider.member.query

星河产品=提供会查询服务
*/
type AlitripMerchantGalaxyProviderMemberQueryAPIResponse struct {
    model.CommonResponse
    AlitripMerchantGalaxyProviderMemberQueryResponse
}

type AlitripMerchantGalaxyProviderMemberQueryResponse struct {
    XMLName xml.Name `xml:"alitrip_merchant_galaxy_provider_member_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 默认描述
    
    Result   *Response `json:"result,omitempty" xml:"result,omitempty"`

    
}
