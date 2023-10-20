package campus

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusDeviceOpenapiSaveeventinfoforibosAPIResponse saveeventinfoforibos API返回值
// alibaba.campus.device.openapi.saveeventinfoforibos
//
// IBos的事件信息上报与反馈的处理接口
type AlibabaCampusDeviceOpenapiSaveeventinfoforibosAPIResponse struct {
	model.CommonResponse
	AlibabaCampusDeviceOpenapiSaveeventinfoforibosAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaCampusDeviceOpenapiSaveeventinfoforibosAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaCampusDeviceOpenapiSaveeventinfoforibosAPIResponseModel).Reset()
}

// AlibabaCampusDeviceOpenapiSaveeventinfoforibosAPIResponseModel is saveeventinfoforibos 成功返回结果
type AlibabaCampusDeviceOpenapiSaveeventinfoforibosAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_campus_device_openapi_saveeventinfoforibos_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// errorMsg
	RequestErrorMsg string `json:"request_error_msg,omitempty" xml:"request_error_msg,omitempty"`
	// success
	RequestSuccess bool `json:"request_success,omitempty" xml:"request_success,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaCampusDeviceOpenapiSaveeventinfoforibosAPIResponseModel) Reset() {
	m.RequestId = ""
	m.RequestErrorMsg = ""
	m.RequestSuccess = false
}

var poolAlibabaCampusDeviceOpenapiSaveeventinfoforibosAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaCampusDeviceOpenapiSaveeventinfoforibosAPIResponse)
	},
}

// GetAlibabaCampusDeviceOpenapiSaveeventinfoforibosAPIResponse 从 sync.Pool 获取 AlibabaCampusDeviceOpenapiSaveeventinfoforibosAPIResponse
func GetAlibabaCampusDeviceOpenapiSaveeventinfoforibosAPIResponse() *AlibabaCampusDeviceOpenapiSaveeventinfoforibosAPIResponse {
	return poolAlibabaCampusDeviceOpenapiSaveeventinfoforibosAPIResponse.Get().(*AlibabaCampusDeviceOpenapiSaveeventinfoforibosAPIResponse)
}

// ReleaseAlibabaCampusDeviceOpenapiSaveeventinfoforibosAPIResponse 将 AlibabaCampusDeviceOpenapiSaveeventinfoforibosAPIResponse 保存到 sync.Pool
func ReleaseAlibabaCampusDeviceOpenapiSaveeventinfoforibosAPIResponse(v *AlibabaCampusDeviceOpenapiSaveeventinfoforibosAPIResponse) {
	v.Reset()
	poolAlibabaCampusDeviceOpenapiSaveeventinfoforibosAPIResponse.Put(v)
}
