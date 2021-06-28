package iot

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
设备控制自定义扩展接口 APIResponse
taobao.ailab.aicloud.top.device.control.custom

设备控制自定义扩展接口
*/
type TaobaoAilabAicloudTopDeviceControlCustomAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"ailab_aicloud_top_device_control_custom_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 业务请求是否成功
    
    Model   bool `json:"model,omitempty" xml:"