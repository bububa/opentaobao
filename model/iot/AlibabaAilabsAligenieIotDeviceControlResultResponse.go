package iot

import (
    "github.com/bububa/opentaobao/model"
)

/* 
设备控制结果 APIResponse
alibaba.ailabs.aligenie.iot.device.control.result

智能IOT解决外部厂商在云云模式在用户通过天猫精灵下发设备指令过程中，厂商指令完成，回调结果通知
*/
type AlibabaAilabsAligenieIotDeviceControlResultAPIResponse struct {
    model.CommonResponse
    Response *AlibabaAilabsAligenieIotDeviceControlResultResponse `json:"alibaba_ailabs_aligenie_iot_device_control_result_response,omitempty"`
}

type AlibabaAilabsAligenieIotDeviceControlResultResponse struct {

    // statusCode
    StatusCode   int64 `json:"status_code,omitempty"`

    // message
    Message   string `json:"message,omitempty"`

    // result
    Result   bool `json:"result,omitempty"`

}
