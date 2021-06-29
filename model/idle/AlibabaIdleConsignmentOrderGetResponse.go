package idle

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
闲鱼帮卖订单查询 APIResponse
alibaba.idle.consignment.order.get

闲鱼帮卖服务商以闲鱼交易买家身份查询订单信息
*/
type AlibabaIdleConsignmentOrderGetAPIResponse struct {
    model.CommonResponse
    AlibabaIdleConsignmentOrderGetResponse
}

type AlibabaIdleConsignmentOrderGetResponse struct {
    XMLName xml.Name `xml:"alibaba_idle_consignment_order_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口返回model
    
    Result   *AlibabaIdleConsignmentOrderGetResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
