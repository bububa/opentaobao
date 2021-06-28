package alink

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
更新设备昵称等信息 APIResponse
alibaba.alink.device.info.update

更新设备昵称等信息
*/
type AlibabaAlinkDeviceInfoUpdateAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_alink_device_info_update_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 结果
    
    Result   *TopServiceResult `json:"result,omitempty" xml:"