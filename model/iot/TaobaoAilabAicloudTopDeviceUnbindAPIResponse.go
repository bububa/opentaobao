package iot

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAilabAicloudTopDeviceUnbindAPIResponse 解绑设备 API返回值
// taobao.ailab.aicloud.top.device.unbind
//
// 解绑设备
type TaobaoAilabAicloudTopDeviceUnbindAPIResponse struct {
	model.CommonResponse
	TaobaoAilabAicloudTopDeviceUnbindAPIResponseModel
}

// TaobaoAilabAicloudTopDeviceUnbindAPIResponseModel is 解绑设备 成功返回结果
type TaobaoAilabAicloudTopDeviceUnbindAPIResponseModel struct {
	XMLName xml.Name `xml:"ailab_aicloud_top_device_unbind_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 解绑是否成功
	Model bool `json:"model,omitempty" xml:"model,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
	// msgInfo错误码信息，成功返回null
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
}
