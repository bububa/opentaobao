package alink

import (
    "github.com/bububa/opentaobao/model"
)

/* 
更新设备昵称等信息 APIResponse
alibaba.alink.device.info.update

更新设备昵称等信息
*/
type AlibabaAlinkDeviceInfoUpdateAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaAlinkDeviceInfoUpdateResponse `json:"alibaba_alink_device_info_update_response,omitempty"` 
    AlibabaAlinkDeviceInfoUpdateResponse
}

/* model for simplify = false
type AlibabaAlinkDeviceInfoUpdateResponse struct {

    // 结果
    
    Result  *struct {
        TopServiceResult  *TopServiceResult `json:"top_service_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaAlinkDeviceInfoUpdateResponse struct {

    // 结果
    Result   *TopServiceResult `json:"result,omitempty"`

}
