package logistic

import (
    "github.com/bububa/opentaobao/model"
)

/* 
物流信息回传接口 APIResponse
alibaba.alihealth.nr.delivery.history.save

商家ERP回传物流信息
*/
type AlibabaAlihealthNrDeliveryHistorySaveAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaAlihealthNrDeliveryHistorySaveResponse `json:"alibaba_alihealth_nr_delivery_history_save_response,omitempty"` 
    AlibabaAlihealthNrDeliveryHistorySaveResponse
}

/* model for simplify = false
type AlibabaAlihealthNrDeliveryHistorySaveResponse struct {

    // 返回结果
    
    Result  *struct {
        ResponseResult  *ResponseResult `json:"response_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaAlihealthNrDeliveryHistorySaveResponse struct {

    // 返回结果
    Result   *ResponseResult `json:"result,omitempty"`

}
