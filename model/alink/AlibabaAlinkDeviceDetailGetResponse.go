package alink

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取设备详情 API返回值 
alibaba.alink.device.detail.get

阿里智能获取设备详情
*/
type AlibabaAlinkDeviceDetailGetAPIResponse struct {
    model.CommonResponse
    AlibabaAlinkDeviceDetailGetResponse
}

// 获取设备详情 成功返回结果
type AlibabaAlinkDeviceDetailGetResponse struct {
    XMLName xml.Name `xml:"alibaba_alink_device_detail_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 结果
    Result   *TopServiceResult `json:"result,omitempty" xml:"result,omitempty"`
}
