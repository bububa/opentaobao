package idle

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
闲鱼寄卖V2订单查询 APIResponse
alibaba.idle.consignmentii.order.get

闲鱼寄卖V2服务商以闲鱼交易买家身份查询订单信息
*/
type AlibabaIdleConsignmentiiOrderGetAPIResponse struct {
    model.CommonResponse
    AlibabaIdleConsignmentiiOrderGetResponse
}

type AlibabaIdleConsignmentiiOrderGetResponse struct {
    XMLName xml.Name `xml:"alibaba_idle_consignmentii_order_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口返回model
    
    Result   *AlibabaIdleConsignmentiiOrderGetResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
