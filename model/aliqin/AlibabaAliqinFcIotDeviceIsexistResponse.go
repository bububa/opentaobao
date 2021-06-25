package aliqin

import (
    "github.com/bububa/opentaobao/model"
)

/* 
判断设备是否存在 APIResponse
alibaba.aliqin.fc.iot.device.isexist

判断设备是否存在
*/
type AlibabaAliqinFcIotDeviceIsexistAPIResponse struct {
    model.CommonResponse
    Response *AlibabaAliqinFcIotDeviceIsexistResponse `json:"alibaba_aliqin_fc_iot_device_isexist_response,omitempty"`
}

type AlibabaAliqinFcIotDeviceIsexistResponse struct {

    // result
    Result   *AlibabaAliqinFcIotDeviceIsexistResult `json:"result,omitempty"`

}
