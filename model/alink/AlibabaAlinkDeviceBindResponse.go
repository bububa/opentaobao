package alink

import (
    "github.com/bububa/opentaobao/model"
)

/* 
绑定设备 APIResponse
alibaba.alink.device.bind

阿里智能解绑设备
*/
type AlibabaAlinkDeviceBindAPIResponse struct {
    model.CommonResponse
    Response *AlibabaAlinkDeviceBindResponse `json:"alibaba_alink_device_bind_response,omitempty"`
}

type AlibabaAlinkDeviceBindResponse struct {

    // 结果
    Result   *TopServiceResult `json:"result,omitempty"`

}
