package aliqin

import (
    "github.com/bububa/opentaobao/model"
)

/* 
商家提交设备信息 APIResponse
alibaba.aliqin.fc.iot.device.post

物联网商家设备信息录入
*/
type AlibabaAliqinFcIotDevicePostAPIResponse struct {
    model.CommonResponse
    Response *AlibabaAliqinFcIotDevicePostResponse `json:"alibaba_aliqin_fc_iot_device_post_response,omitempty"`
}

type AlibabaAliqinFcIotDevicePostResponse struct {

    // result
    Result   *AlibabaAliqinFcIotDevicePostResult `json:"result,omitempty"`

}
