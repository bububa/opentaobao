package idle

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
闲鱼无需物流发货 APIResponse
alibaba.idle.order.dummy.send

适用于电子卡券等虚拟商品不需要物流的商品发货。
*/
type AlibabaIdleOrderDummySendAPIResponse struct {
    model.CommonResponse
    AlibabaIdleOrderDummySendResponse
}

type AlibabaIdleOrderDummySendResponse struct {
    XMLName xml.Name `xml:"alibaba_idle_order_dummy_send_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回包装类
    
    Result   *BaseResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
