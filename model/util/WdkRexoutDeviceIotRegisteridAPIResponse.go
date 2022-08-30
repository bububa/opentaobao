package util

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// WdkRexoutDeviceIotRegisteridAPIResponse 通过设备ID获取三元组-外部 API返回值
// wdk.rexout.device.iot.registerid
//
// 通过设备ID获取三元组-外部
type WdkRexoutDeviceIotRegisteridAPIResponse struct {
	model.CommonResponse
	WdkRexoutDeviceIotRegisteridAPIResponseModel
}

// WdkRexoutDeviceIotRegisteridAPIResponseModel is 通过设备ID获取三元组-外部 成功返回结果
type WdkRexoutDeviceIotRegisteridAPIResponseModel struct {
	XMLName xml.Name `xml:"wdk_rexout_device_iot_registerid_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误信息
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 错误码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 数据
	Data *IotRegisterResult `json:"data,omitempty" xml:"data,omitempty"`
	// 结果
	Succeed bool `json:"succeed,omitempty" xml:"succeed,omitempty"`
}
