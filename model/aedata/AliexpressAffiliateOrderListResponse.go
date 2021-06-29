package aedata

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
AE推广者订单批量获取接口 APIResponse
aliexpress.affiliate.order.list

AE联盟推广者订单分页查询接口
*/
type AliexpressAffiliateOrderListAPIResponse struct {
    model.CommonResponse
    AliexpressAffiliateOrderListResponse
}

type AliexpressAffiliateOrderListResponse struct {
    XMLName xml.Name `xml:"aliexpress_affiliate_order_list_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    RespResult   *ResponseDto `json:"resp_result,omitempty" xml:"resp_result,omitempty"`

    
}
