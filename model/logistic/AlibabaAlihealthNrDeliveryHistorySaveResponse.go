package logistic

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
物流信息回传接口 APIResponse
alibaba.alihealth.nr.delivery.history.save

商家ERP回传物流信息
*/
type AlibabaAlihealthNrDeliveryHistorySaveAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_alihealth_nr_delivery_history_save_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回结果
    
    Result   *ResponseResult `json:"result,omitempty" xml:"