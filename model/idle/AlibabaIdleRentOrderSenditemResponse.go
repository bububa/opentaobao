package idle

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
确认发货 APIResponse
alibaba.idle.rent.order.senditem

确认发货
*/
type AlibabaIdleRentOrderSenditemAPIResponse struct {
    model.CommonResponse
    AlibabaIdleRentOrderSenditemResponse
}

type AlibabaIdleRentOrderSenditemResponse struct {
    XMLName xml.Name `xml:"alibaba_idle_rent_order_senditem_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 系统自动生成
    
    Result   *TopResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
