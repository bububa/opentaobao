package alihealth2

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
ISV解绑商品 APIResponse
alibaba.alihealth.dental.item.unbind

ISV解绑商品
*/
type AlibabaAlihealthDentalItemUnbindAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthDentalItemUnbindResponse
}

type AlibabaAlihealthDentalItemUnbindResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_dental_item_unbind_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *AlibabaAlihealthDentalItemUnbindMtopResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
