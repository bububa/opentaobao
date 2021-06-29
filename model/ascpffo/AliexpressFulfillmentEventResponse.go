package ascpffo

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
AE履约事件处理 APIResponse
aliexpress.fulfillment.event

AE用 履约底层声明发货能力
*/
type AliexpressFulfillmentEventAPIResponse struct {
    model.CommonResponse
    AliexpressFulfillmentEventResponse
}

type AliexpressFulfillmentEventResponse struct {
    XMLName xml.Name `xml:"aliexpress_fulfillment_event_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口返回model
    
    Result   *AliexpressFulfillmentEventResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
