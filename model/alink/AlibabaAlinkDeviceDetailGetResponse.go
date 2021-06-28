package alink

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取设备详情 APIResponse
alibaba.alink.device.detail.get

阿里智能获取设备详情
*/
type AlibabaAlinkDeviceDetailGetAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaAlinkDeviceDetailGetResponse `json:"alibaba_alink_device_detail_get_response,omitempty"` 
    AlibabaAlinkDeviceDetailGetResponse
}

/* model for simplify = false
type AlibabaAlinkDeviceDetailGetResponse struct {

    // 结果
    
    Result  *struct {
        TopServiceResult  *TopServiceResult `json:"top_service_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaAlinkDeviceDetailGetResponse struct {

    // 结果
    Result   *TopServiceResult `json:"result,omitempty"`

}
