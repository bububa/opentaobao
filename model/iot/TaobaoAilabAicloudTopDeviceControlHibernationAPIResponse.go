package iot

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAilabAicloudTopDeviceControlHibernationAPIResponse 定时休眠 API返回值
// taobao.ailab.aicloud.top.device.control.hibernation
//
// 定时休眠
type TaobaoAilabAicloudTopDeviceControlHibernationAPIResponse struct {
	model.CommonResponse
	TaobaoAilabAicloudTopDeviceControlHibernationAPIResponseModel
}

// TaobaoAilabAicloudTopDeviceControlHibernationAPIResponseModel is 定时休眠 成功返回结果
type TaobaoAilabAicloudTopDeviceControlHibernationAPIResponseModel struct {
	XMLName xml.Name `xml:"ailab_aicloud_top_device_control_hibernation_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// msgCode
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// msgInfo
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 业务请求是否成功
	Model bool `json:"model,omitempty" xml:"model,omitempty"`
	// 网络请求是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
