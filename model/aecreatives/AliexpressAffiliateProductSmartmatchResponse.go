package aecreatives

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
联盟物料智能推荐api API返回值 
aliexpress.affiliate.product.smartmatch

联盟物料算法智能推荐
*/
type AliexpressAffiliateProductSmartmatchAPIResponse struct {
    model.CommonResponse
    AliexpressAffiliateProductSmartmatchResponse
}

// 联盟物料智能推荐api 成功返回结果
type AliexpressAffiliateProductSmartmatchResponse struct {
    XMLName xml.Name `xml:"aliexpress_affiliate_product_smartmatch_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回结果
    RespResult   *ResponseDTO `json:"resp_result,omitempty" xml:"resp_result,omitempty"`
}
