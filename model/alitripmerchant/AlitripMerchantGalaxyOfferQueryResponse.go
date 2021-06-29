package alitripmerchant

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
星河-offer查询 APIResponse
alitrip.merchant.galaxy.offer.query

根据offer的ID查询offer信息
*/
type AlitripMerchantGalaxyOfferQueryAPIResponse struct {
    model.CommonResponse
    AlitripMerchantGalaxyOfferQueryResponse
}

type AlitripMerchantGalaxyOfferQueryResponse struct {
    XMLName xml.Name `xml:"alitrip_merchant_galaxy_offer_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 默认描述
    
    Result   *Response `json:"result,omitempty" xml:"result,omitempty"`

    
}
