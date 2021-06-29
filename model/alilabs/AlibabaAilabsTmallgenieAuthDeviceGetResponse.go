package alilabs

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取设备详情 API返回值 
alibaba.ailabs.tmallgenie.auth.device.get

通过此接口获取设备详情
*/
type AlibabaAilabsTmallgenieAuthDeviceGetAPIResponse struct {
    model.CommonResponse
    AlibabaAilabsTmallgenieAuthDeviceGetResponse
}

// 获取设备详情 成功返回结果
type AlibabaAilabsTmallgenieAuthDeviceGetResponse struct {
    XMLName xml.Name `xml:"alibaba_ailabs_tmallgenie_auth_device_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 设备信息
    DeviceInfo   *AlibabaAilabsTmallgenieAuthDeviceGetResult `json:"device_info,omitempty" xml:"device_info,omitempty"`
}
