package alink

import (
    "github.com/bububa/opentaobao/model"
)

/* 
解绑设备 APIResponse
alibaba.alink.device.unbind

阿里智能解绑设备
*/
type AlibabaAlinkDeviceUnbindAPIResponse struct {
    model.CommonResponse
    Response *AlibabaAlinkDeviceUnbindResponse `json:"alibaba_alink_device_unbind_response,omitempty"`
}

type AlibabaAlinkDeviceUnbindResponse struct {

    // 结果
    Result   *TopServiceResult `json:"result,omitempty"`

}
