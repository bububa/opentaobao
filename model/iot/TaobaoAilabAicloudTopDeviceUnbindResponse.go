package iot

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
解绑设备 APIResponse
taobao.ailab.aicloud.top.device.unbind

解绑设备
*/
type TaobaoAilabAicloudTopDeviceUnbindAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"ailab_aicloud_top_device_unbind_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 解绑是否成功
    
    Model   bool `json:"model,omitempty" xml:"