package idle

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
创建揽收物流 APIResponse
alibaba.idle.rent.order.logistics.deliver

创建揽收物流
商家去物流公司创建物流订单
*/
type AlibabaIdleRentOrderLogisticsDeliverAPIResponse struct {
    model.CommonResponse
    AlibabaIdleRentOrderLogisticsDeliverResponse
}

type AlibabaIdleRentOrderLogisticsDeliverResponse struct {
    XMLName xml.Name `xml:"alibaba_idle_rent_order_logistics_deliver_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 系统自动生成
    
    Result   *TopResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
