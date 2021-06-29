package drug

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
上传订单同城快递单号 APIResponse
alibaba.health.nr.logistics.deliveryno.update

上传订单同城快递单号
*/
type AlibabaHealthNrLogisticsDeliverynoUpdateAPIResponse struct {
    model.CommonResponse
    AlibabaHealthNrLogisticsDeliverynoUpdateResponse
}

type AlibabaHealthNrLogisticsDeliverynoUpdateResponse struct {
    XMLName xml.Name `xml:"alibaba_health_nr_logistics_deliveryno_update_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 更新成功
    
    Result   bool `json:"result,omitempty" xml:"result,omitempty"`

    
}
