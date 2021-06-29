package ottpay

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
iot设备列表变化接口 APIResponse
youku.ott.iot.devicelist.change

iot设备列表变化接口
*/
type YoukuOttIotDevicelistChangeAPIResponse struct {
    model.CommonResponse
    YoukuOttIotDevicelistChangeResponse
}

type YoukuOttIotDevicelistChangeResponse struct {
    XMLName xml.Name `xml:"youku_ott_iot_devicelist_change_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // msgInfo
    
    MsgInfo   string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`

    
    // 成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
}
