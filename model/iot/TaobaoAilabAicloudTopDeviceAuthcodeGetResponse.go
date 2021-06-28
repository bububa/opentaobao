package iot

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取设备授权码 APIResponse
taobao.ailab.aicloud.top.device.authcode.get

获取设备授权码
*/
type TaobaoAilabAicloudTopDeviceAuthcodeGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"ailab_aicloud_top_device_authcode_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 设备授权码，后续流程中所述的auth code
    
    Model   string `json:"model,omitempty" xml:"