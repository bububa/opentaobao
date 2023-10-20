package tmallgenie

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAilabsIotDeviceMeshEventInvokeAPIResponse 弹内设备中心事件调用 API返回值
// alibaba.ailabs.iot.device.mesh.event.invoke
//
// 弹内设备中心事件调用
type AlibabaAilabsIotDeviceMeshEventInvokeAPIResponse struct {
	model.CommonResponse
	AlibabaAilabsIotDeviceMeshEventInvokeAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAilabsIotDeviceMeshEventInvokeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAilabsIotDeviceMeshEventInvokeAPIResponseModel).Reset()
}

// AlibabaAilabsIotDeviceMeshEventInvokeAPIResponseModel is 弹内设备中心事件调用 成功返回结果
type AlibabaAilabsIotDeviceMeshEventInvokeAPIResponseModel struct {
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

// Reset 清空结构体
func (m *AlibabaAilabsIotDeviceMeshEventInvokeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.TraceId = ""
	m.RetValue = ""
	m.Message = ""
	m.RetCode = 0
	m.Result = false
}

var poolAlibabaAilabsIotDeviceMeshEventInvokeAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAilabsIotDeviceMeshEventInvokeAPIResponse)
	},
}

// GetAlibabaAilabsIotDeviceMeshEventInvokeAPIResponse 从 sync.Pool 获取 AlibabaAilabsIotDeviceMeshEventInvokeAPIResponse
func GetAlibabaAilabsIotDeviceMeshEventInvokeAPIResponse() *AlibabaAilabsIotDeviceMeshEventInvokeAPIResponse {
	return poolAlibabaAilabsIotDeviceMeshEventInvokeAPIResponse.Get().(*AlibabaAilabsIotDeviceMeshEventInvokeAPIResponse)
}

// ReleaseAlibabaAilabsIotDeviceMeshEventInvokeAPIResponse 将 AlibabaAilabsIotDeviceMeshEventInvokeAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAilabsIotDeviceMeshEventInvokeAPIResponse(v *AlibabaAilabsIotDeviceMeshEventInvokeAPIResponse) {
	v.Reset()
	poolAlibabaAilabsIotDeviceMeshEventInvokeAPIResponse.Put(v)
}
