package aecreatives

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
联盟推广商品获取接口 API返回值 
aliexpress.affiliate.product.query

联盟推广商品搜索接口，用于搜索联盟推广商品数据
*/
type AliexpressAffiliateProductQueryAPIResponse struct {
    model.CommonResponse
    AliexpressAffiliateProductQueryAPIResponseModel
}

// 联盟推广商品获取接口 成功返回结果
type AliexpressAffiliateProductQueryAPIResponseModel struct {
    XMLName xml.Name `xml:"aliexpress_affiliate_product_query_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回结果
    RespResult   *ResponseDto `json:"resp_result,omitempty" xml:"resp_result,omitempty"`
}
