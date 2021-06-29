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
    TaobaoAilabAicloudTopDeviceUnbindResponse
}

type TaobaoAilabAicloudTopDeviceUnbindResponse struct {
    XMLName xml.Name `xml:"ailab_aicloud_top_device_unbind_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 解绑是否成功
    
    Model   bool `json:"model,omitempty" xml:"model,omitempty"`

    
    // 是否成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
    // msgInfo错误码信息，成功返回null
    
    MsgInfo   string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`

    
}
