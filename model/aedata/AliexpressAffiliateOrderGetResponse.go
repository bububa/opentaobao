package aedata

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
AE流量订单详情获取API API返回值 
aliexpress.affiliate.order.get

联盟推广订单效果获取API
*/
type AliexpressAffiliateOrderGetAPIResponse struct {
    model.CommonResponse
    AliexpressAffiliateOrderGetResponse
}

// AE流量订单详情获取API 成功返回结果
type AliexpressAffiliateOrderGetResponse struct {
    XMLName xml.Name `xml:"aliexpress_affiliate_order_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回结果
    RespResult   *ResponseDTO `json:"resp_result,omitempty" xml:"resp_result,omitempty"`
}
