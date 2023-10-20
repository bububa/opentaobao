package campus

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusDeviceOpenapiGettemplatelistAPIResponse 查询设备模板 API返回值
// alibaba.campus.device.openapi.gettemplatelist
//
// 查询设备模板信息
type AlibabaCampusDeviceOpenapiGettemplatelistAPIResponse struct {
	model.CommonResponse
	AlibabaCampusDeviceOpenapiGettemplatelistAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaCampusDeviceOpenapiGettemplatelistAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaCampusDeviceOpenapiGettemplatelistAPIResponseModel).Reset()
}

// AlibabaCampusDeviceOpenapiGettemplatelistAPIResponseModel is 查询设备模板 成功返回结果
type AlibabaCampusDeviceOpenapiGettemplatelistAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_campus_device_openapi_gettemplatelist_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *PageResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaCampusDeviceOpenapiGettemplatelistAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaCampusDeviceOpenapiGettemplatelistAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaCampusDeviceOpenapiGettemplatelistAPIResponse)
	},
}

// GetAlibabaCampusDeviceOpenapiGettemplatelistAPIResponse 从 sync.Pool 获取 AlibabaCampusDeviceOpenapiGettemplatelistAPIResponse
func GetAlibabaCampusDeviceOpenapiGettemplatelistAPIResponse() *AlibabaCampusDeviceOpenapiGettemplatelistAPIResponse {
	return poolAlibabaCampusDeviceOpenapiGettemplatelistAPIResponse.Get().(*AlibabaCampusDeviceOpenapiGettemplatelistAPIResponse)
}

// ReleaseAlibabaCampusDeviceOpenapiGettemplatelistAPIResponse 将 AlibabaCampusDeviceOpenapiGettemplatelistAPIResponse 保存到 sync.Pool
func ReleaseAlibabaCampusDeviceOpenapiGettemplatelistAPIResponse(v *AlibabaCampusDeviceOpenapiGettemplatelistAPIResponse) {
	v.Reset()
	poolAlibabaCampusDeviceOpenapiGettemplatelistAPIResponse.Put(v)
}
