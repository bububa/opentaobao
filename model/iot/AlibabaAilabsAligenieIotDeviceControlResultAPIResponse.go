package iot

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAilabsAligenieIotDeviceControlResultAPIResponse
设备控制结果 API返回值
alibaba.ailabs.aligenie.iot.device.control.result

智能IOT解决外部厂商在云云模式在用户通过天猫精灵下发设备指令过程中，厂商指令完成，回调结果通知 */
type AlibabaAilabsAligenieIotDeviceControlResultAPIResponse struct {
	model.CommonResponse
	AlibabaAilabsAligenieIotDeviceControlResultAPIResponseModel
}

// AlibabaAilabsAligenieIotDeviceControlResultAPIResponseModel is 设备控制结果 成功返回结果
type AlibabaAilabsAligenieIotDeviceControlResultAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ailabs_aligenie_iot_device_control_result_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// statusCode
	StatusCode int64 `json:"status_code,omitempty" xml:"status_code,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// result
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}
