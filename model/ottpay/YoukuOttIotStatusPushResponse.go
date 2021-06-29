package ottpay

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
iot设备状态变化通知接口 APIResponse
youku.ott.iot.status.push

ott iot设备状态通知
*/
type YoukuOttIotStatusPushAPIResponse struct {
    model.CommonResponse
    YoukuOttIotStatusPushResponse
}

type YoukuOttIotStatusPushResponse struct {
    XMLName xml.Name `xml:"youku_ott_iot_status_push_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // msgInfo
    
    MsgInfo   string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`

    
    // 成功标识
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
}
