package alink

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取设备详情 APIResponse
alibaba.alink.device.detail.get

阿里智能获取设备详情
*/
type AlibabaAlinkDeviceDetailGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_alink_device_detail_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 结果
    
    Result   *TopServiceResult `json:"result,omitempty" xml:"