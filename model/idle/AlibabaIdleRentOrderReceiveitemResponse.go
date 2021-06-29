package idle

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
确认签收 APIResponse
alibaba.idle.rent.order.receiveitem

确认揽收/签收
*/
type AlibabaIdleRentOrderReceiveitemAPIResponse struct {
    model.CommonResponse
    AlibabaIdleRentOrderReceiveitemResponse
}

type AlibabaIdleRentOrderReceiveitemResponse struct {
    XMLName xml.Name `xml:"alibaba_idle_rent_order_receiveitem_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 系统自动生成
    
    Result   *TopResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
