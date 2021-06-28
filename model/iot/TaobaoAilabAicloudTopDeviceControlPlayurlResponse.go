package iot

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
点播url APIResponse
taobao.ailab.aicloud.top.device.control.playurl

点播url
*/
type TaobaoAilabAicloudTopDeviceControlPlayurlAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"ailab_aicloud_top_device_control_playurl_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 业务请求是否成功
    
    Model   bool `json:"model,omitempty" xml:"