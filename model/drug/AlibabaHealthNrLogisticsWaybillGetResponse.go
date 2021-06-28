package drug

import (
    "github.com/bububa/opentaobao/model"
)

/* 
电子面单查询接口 APIResponse
alibaba.health.nr.logistics.waybill.get

商家登录后根据订单号查询物流单号及电子面单信息
*/
type AlibabaHealthNrLogisticsWaybillGetAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaHealthNrLogisticsWaybillGetResponse `json:"alibaba_health_nr_logistics_waybill_get_response,omitempty"` 
    AlibabaHealthNrLogisticsWaybillGetResponse
}

/* model for simplify = false
type AlibabaHealthNrLogisticsWaybillGetResponse struct {

    // 响应结果对象
    
    ResponseResult  *struct {
        ResponseResult  *ResponseResult `json:"response_result,omitempty"`
    } `json:"response_result,omitempty"`
    

}
*/

type AlibabaHealthNrLogisticsWaybillGetResponse struct {

    // 响应结果对象
    ResponseResult   *ResponseResult `json:"response_result,omitempty"`

}
