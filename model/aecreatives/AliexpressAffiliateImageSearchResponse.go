package aecreatives

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
图搜 API返回值 
aliexpress.affiliate.image.search

图片搜索接口
*/
type AliexpressAffiliateImageSearchAPIResponse struct {
    model.CommonResponse
    AliexpressAffiliateImageSearchResponse
}

// 图搜 成功返回结果
type AliexpressAffiliateImageSearchResponse struct {
    XMLName xml.Name `xml:"aliexpress_affiliate_image_search_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 默认描述
    Result   *Response `json:"result,omitempty" xml:"result,omitempty"`
}
