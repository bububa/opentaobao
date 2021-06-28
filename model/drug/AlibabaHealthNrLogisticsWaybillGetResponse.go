package drug

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
电子面单查询接口 APIResponse
alibaba.health.nr.logistics.waybill.get

商家登录后根据订单号查询物流单号及电子面单信息
*/
type AlibabaHealthNrLogisticsWaybillGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_health_nr_logistics_waybill_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 响应结果对象
    
    ResponseResult   *ResponseResult `json:"response_result,omitempty" xml:"