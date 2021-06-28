package drug

import (
    "github.com/bububa/opentaobao/model"
)

/* 
上传订单同城快递单号 APIResponse
alibaba.health.nr.logistics.deliveryno.update

上传订单同城快递单号
*/
type AlibabaHealthNrLogisticsDeliverynoUpdateAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaHealthNrLogisticsDeliverynoUpdateResponse `json:"alibaba_health_nr_logistics_deliveryno_update_response,omitempty"` 
    AlibabaHealthNrLogisticsDeliverynoUpdateResponse
}

/* model for simplify = false
type AlibabaHealthNrLogisticsDeliverynoUpdateResponse struct {

    // 更新成功
    
    Result   bool `json:"result,omitempty"`
    

}
*/

type AlibabaHealthNrLogisticsDeliverynoUpdateResponse struct {

    // 更新成功
    Result   bool `json:"result,omitempty"`

}
