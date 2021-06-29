package lstlogistics2

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
供应商-异云-无需物流发货 APIResponse
alibaba.lst.logistics.notrace.send

异地云仓的订单，使用无需物流的发货方式来变更订单发货状态
*/
type AlibabaLstLogisticsNotraceSendAPIResponse struct {
    model.CommonResponse
    AlibabaLstLogisticsNotraceSendResponse
}

type AlibabaLstLogisticsNotraceSendResponse struct {
    XMLName xml.Name `xml:"alibaba_lst_logistics_notrace_send_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口返回model
    
    Result   *AlibabaLstLogisticsNotraceSendResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
