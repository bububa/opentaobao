package aecreatives

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
图搜 APIResponse
aliexpress.affiliate.image.search

图片搜索接口
*/
type AliexpressAffiliateImageSearchAPIResponse struct {
    model.CommonResponse
    AliexpressAffiliateImageSearchResponse
}

type AliexpressAffiliateImageSearchResponse struct {
    XMLName xml.Name `xml:"aliexpress_affiliate_image_search_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 默认描述
    
    Result   *Response `json:"result,omitempty" xml:"result,omitempty"`

    
}
