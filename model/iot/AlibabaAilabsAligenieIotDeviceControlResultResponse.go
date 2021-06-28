package iot

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
设备控制结果 APIResponse
alibaba.ailabs.aligenie.iot.device.control.result

智能IOT解决外部厂商在云云模式在用户通过天猫精灵下发设备指令过程中，厂商指令完成，回调结果通知
*/
type AlibabaAilabsAligenieIotDeviceControlResultAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_ailabs_aligenie_iot_device_control_result_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // statusCode
    
    StatusCode   int64 `json:"status_code,omitempty" xml:"