package tmallgenie

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaailabsiotdevicemesheventinvokeAPIResponse 弹内设备中心事件调用 API返回值
// alibaba.ailabs.iot.device.mesh.event.invoke
//
// 弹内设备中心事件调用
type AlibabaailabsiotdevicemesheventinvokeAPIResponse struct {
	model.CommonResponse
	AlibabaailabsiotdevicemesheventinvokeAPIResponseModel
}

// AlibabaailabsiotdevicemesheventinvokeAPIResponseModel is 弹内设备中心事件调用 成功返回结果
type AlibabaailabsiotdevicemesheventinvokeAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ailabs_iot_device_mesh_event_invoke_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 事件链路id
	TraceId string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
	// 响应值
	RetValue string `json:"ret_value,omitempty" xml:"ret_value,omitempty"`
	// 错误描述
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 错误码
	RetCode int64 `json:"ret_code,omitempty" xml:"ret_code,omitempty"`
	// 成功还是失败
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}
