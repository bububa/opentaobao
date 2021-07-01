package alilabs

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
ailabs iot 设备状态更新 API返回值 
alibaba.ailabs.iot.device.status.update

用于人工智能实验室IoT合作厂商上报三方接入IoT设备状态更新时的设备状态上报
*/
type AlibabaAilabsIotDeviceStatusUpdateAPIResponse struct {
    model.CommonResponse
    AlibabaAilabsIotDeviceStatusUpdateAPIResponseModel
}

// ailabs iot 设备状态更新 成功返回结果
type AlibabaAilabsIotDeviceStatusUpdateAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_ailabs_iot_device_status_update_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 设备状态更新是否成功
    Result   *AlibabaAilabsIotDeviceStatusUpdateResult `json:"result,omitempty" xml:"result,omitempty"`
}
