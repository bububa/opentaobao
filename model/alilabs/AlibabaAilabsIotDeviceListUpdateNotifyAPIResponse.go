package alilabs

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
设备列表更新通知 API返回值 
alibaba.ailabs.iot.device.list.update.notify

用于人工智能实验室IoT合作厂商上报三方接入IoT设备列表更新时的通知
*/
type AlibabaAilabsIotDeviceListUpdateNotifyAPIResponse struct {
    model.CommonResponse
    AlibabaAilabsIotDeviceListUpdateNotifyAPIResponseModel
}

// 设备列表更新通知 成功返回结果
type AlibabaAilabsIotDeviceListUpdateNotifyAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_ailabs_iot_device_list_update_notify_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 接口返回model
    Result   *AlibabaAilabsIotDeviceListUpdateNotifyResult `json:"result,omitempty" xml:"result,omitempty"`
}
