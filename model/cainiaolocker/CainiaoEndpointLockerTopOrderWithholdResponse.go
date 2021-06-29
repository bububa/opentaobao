package cainiaolocker

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
代扣支付 APIResponse
cainiao.endpoint.locker.top.order.withhold

提供代扣，允许有一笔欠款。
*/
type CainiaoEndpointLockerTopOrderWithholdAPIResponse struct {
    model.CommonResponse
    CainiaoEndpointLockerTopOrderWithholdResponse
}

type CainiaoEndpointLockerTopOrderWithholdResponse struct {
    XMLName xml.Name `xml:"cainiao_endpoint_locker_top_order_withhold_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *SingleResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
