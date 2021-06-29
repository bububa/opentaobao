package idle

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
闲鱼回收订单查询V1.1 APIResponse
alibaba.idle.recycle.order.show

查询回收订单
*/
type AlibabaIdleRecycleOrderShowAPIResponse struct {
    model.CommonResponse
    AlibabaIdleRecycleOrderShowResponse
}

type AlibabaIdleRecycleOrderShowResponse struct {
    XMLName xml.Name `xml:"alibaba_idle_recycle_order_show_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *AlibabaIdleRecycleOrderShowResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
