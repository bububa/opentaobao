package iot

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
设备儿童锁 APIResponse
taobao.ailab.aicloud.top.device.control.childlock

设备儿童锁
*/
type TaobaoAilabAicloudTopDeviceControlChildlockAPIResponse struct {
    model.CommonResponse
    TaobaoAilabAicloudTopDeviceControlChildlockResponse
}

type TaobaoAilabAicloudTopDeviceControlChildlockResponse struct {
    XMLName xml.Name `xml:"ailab_aicloud_top_device_control_childlock_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 业务请求是否成功
    
    Model   bool `json:"model,omitempty" xml:"model,omitempty"`

    
    // msgCode
    
    MsgCode   string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`

    
    // msgInfo
    
    MsgInfo   string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`

    
    // 网络请求是否成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
}
