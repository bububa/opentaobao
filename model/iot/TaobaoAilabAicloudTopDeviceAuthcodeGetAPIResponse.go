package iot

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取设备授权码 API返回值 
taobao.ailab.aicloud.top.device.authcode.get

获取设备授权码
*/
type TaobaoAilabAicloudTopDeviceAuthcodeGetAPIResponse struct {
    model.CommonResponse
    TaobaoAilabAicloudTopDeviceAuthcodeGetAPIResponseModel
}

// 获取设备授权码 成功返回结果
type TaobaoAilabAicloudTopDeviceAuthcodeGetAPIResponseModel struct {
    XMLName xml.Name `xml:"ailab_aicloud_top_device_authcode_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 设备授权码，后续流程中所述的auth code
    Model   string `json:"model,omitempty" xml:"model,omitempty"`
    // 是否成功
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
    // msgInfo错误码信息，成功返回null
    MsgInfo   string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
}
