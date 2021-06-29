package aecreatives

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询联盟爆品数据 APIResponse
aliexpress.affiliate.hotproduct.query

查询联盟爆品API
*/
type AliexpressAffiliateHotproductQueryAPIResponse struct {
    model.CommonResponse
    AliexpressAffiliateHotproductQueryResponse
}

type AliexpressAffiliateHotproductQueryResponse struct {
    XMLName xml.Name `xml:"aliexpress_affiliate_hotproduct_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    RespResult   *ResponseDto `json:"resp_result,omitempty" xml:"resp_result,omitempty"`

    
}
