package bus

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
汽车票退款申请接口 APIResponse
taobao.bus.refundticketprice.set

汽车票代理商利用该接口申请退票
*/
type TaobaoBusRefundticketpriceSetAPIResponse struct {
    model.CommonResponse
    TaobaoBusRefundticketpriceSetResponse
}

type TaobaoBusRefundticketpriceSetResponse struct {
    XMLName xml.Name `xml:"bus_refundticketprice_set_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 退票成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
}
