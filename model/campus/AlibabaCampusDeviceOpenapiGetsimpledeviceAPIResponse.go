package campus

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusDeviceOpenapiGetsimpledeviceAPIResponse 获取单个设备的基础信息(仅包含设备id,code,是否启用,位置信息,描述等基础信息) API返回值
// alibaba.campus.device.openapi.getsimpledevice
//
// 获取指定设备的基础信息(仅包含设备id,code,是否启用,位置信息,描述等基础信息)
type AlibabaCampusDeviceOpenapiGetsimpledeviceAPIResponse struct {
	model.CommonResponse
	AlibabaCampusDeviceOpenapiGetsimpledeviceAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaCampusDeviceOpenapiGetsimpledeviceAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaCampusDeviceOpenapiGetsimpledeviceAPIResponseModel).Reset()
}

// AlibabaCampusDeviceOpenapiGetsimpledeviceAPIResponseModel is 获取单个设备的基础信息(仅包含设备id,code,是否启用,位置信息,描述等基础信息) 成功返回结果
type AlibabaCampusDeviceOpenapiGetsimpledeviceAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_campus_device_openapi_getsimpledevice_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回对象
	Result *PojoResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaCampusDeviceOpenapiGetsimpledeviceAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaCampusDeviceOpenapiGetsimpledeviceAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaCampusDeviceOpenapiGetsimpledeviceAPIResponse)
	},
}

// GetAlibabaCampusDeviceOpenapiGetsimpledeviceAPIResponse 从 sync.Pool 获取 AlibabaCampusDeviceOpenapiGetsimpledeviceAPIResponse
func GetAlibabaCampusDeviceOpenapiGetsimpledeviceAPIResponse() *AlibabaCampusDeviceOpenapiGetsimpledeviceAPIResponse {
	return poolAlibabaCampusDeviceOpenapiGetsimpledeviceAPIResponse.Get().(*AlibabaCampusDeviceOpenapiGetsimpledeviceAPIResponse)
}

// ReleaseAlibabaCampusDeviceOpenapiGetsimpledeviceAPIResponse 将 AlibabaCampusDeviceOpenapiGetsimpledeviceAPIResponse 保存到 sync.Pool
func ReleaseAlibabaCampusDeviceOpenapiGetsimpledeviceAPIResponse(v *AlibabaCampusDeviceOpenapiGetsimpledeviceAPIResponse) {
	v.Reset()
	poolAlibabaCampusDeviceOpenapiGetsimpledeviceAPIResponse.Put(v)
}
