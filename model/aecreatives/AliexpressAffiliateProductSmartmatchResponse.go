package aecreatives

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
联盟物料智能推荐api APIResponse
aliexpress.affiliate.product.smartmatch

联盟物料算法智能推荐
*/
type AliexpressAffiliateProductSmartmatchAPIResponse struct {
    model.CommonResponse
    AliexpressAffiliateProductSmartmatchResponse
}

type AliexpressAffiliateProductSmartmatchResponse struct {
    XMLName xml.Name `xml:"aliexpress_affiliate_product_smartmatch_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    RespResult   *ResponseDto `json:"resp_result,omitempty" xml:"resp_result,omitempty"`

    
}
