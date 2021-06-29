package moscm

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
发货 APIResponse
alibaba.mos.delivery.send

订单发货填写快递单
*/
type AlibabaMosDeliverySendAPIResponse struct {
    model.CommonResponse
    AlibabaMosDeliverySendResponse
}

type AlibabaMosDeliverySendResponse struct {
    XMLName xml.Name `xml:"alibaba_mos_delivery_send_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    Result   *AlibabaMosDeliverySendResultDo `json:"result,omitempty" xml:"result,omitempty"`

    
}
