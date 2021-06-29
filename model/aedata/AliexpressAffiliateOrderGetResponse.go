package aedata

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
AE流量订单详情获取API APIResponse
aliexpress.affiliate.order.get

联盟推广订单效果获取API
*/
type AliexpressAffiliateOrderGetAPIResponse struct {
    model.CommonResponse
    AliexpressAffiliateOrderGetResponse
}

type AliexpressAffiliateOrderGetResponse struct {
    XMLName xml.Name `xml:"aliexpress_affiliate_order_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    RespResult   *ResponseDto `json:"resp_result,omitempty" xml:"resp_result,omitempty"`

    
}
