package iot

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoailabaicloudtopdevicecontrolcustomAPIResponse 设备控制自定义扩展接口 API返回值
// taobao.ailab.aicloud.top.device.control.custom
//
// 设备控制自定义扩展接口
type TaobaoailabaicloudtopdevicecontrolcustomAPIResponse struct {
	model.CommonResponse
	TaobaoailabaicloudtopdevicecontrolcustomAPIResponseModel
}

// TaobaoailabaicloudtopdevicecontrolcustomAPIResponseModel is 设备控制自定义扩展接口 成功返回结果
type TaobaoailabaicloudtopdevicecontrolcustomAPIResponseModel struct {
	XMLName xml.Name `xml:"ailab_aicloud_top_device_control_custom_response"`
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
