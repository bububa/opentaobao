package campus

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusDeviceOpenapiGetsimpledevicelistAPIResponse 查询设备基础信息集合(仅包含设备id,code,是否启用,位置信息,描述等基础信息) API返回值
// alibaba.campus.device.openapi.getsimpledevicelist
//
// 查询设备基础信息集合(仅包含设备id,code,是否启用,位置信息,描述等基础信息)
type AlibabaCampusDeviceOpenapiGetsimpledevicelistAPIResponse struct {
	model.CommonResponse
	AlibabaCampusDeviceOpenapiGetsimpledevicelistAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaCampusDeviceOpenapiGetsimpledevicelistAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaCampusDeviceOpenapiGetsimpledevicelistAPIResponseModel).Reset()
}

// AlibabaCampusDeviceOpenapiGetsimpledevicelistAPIResponseModel is 查询设备基础信息集合(仅包含设备id,code,是否启用,位置信息,描述等基础信息) 成功返回结果
type AlibabaCampusDeviceOpenapiGetsimpledevicelistAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_campus_device_openapi_getsimpledevicelist_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回对象
	Result *PageResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaCampusDeviceOpenapiGetsimpledevicelistAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaCampusDeviceOpenapiGetsimpledevicelistAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaCampusDeviceOpenapiGetsimpledevicelistAPIResponse)
	},
}

// GetAlibabaCampusDeviceOpenapiGetsimpledevicelistAPIResponse 从 sync.Pool 获取 AlibabaCampusDeviceOpenapiGetsimpledevicelistAPIResponse
func GetAlibabaCampusDeviceOpenapiGetsimpledevicelistAPIResponse() *AlibabaCampusDeviceOpenapiGetsimpledevicelistAPIResponse {
	return poolAlibabaCampusDeviceOpenapiGetsimpledevicelistAPIResponse.Get().(*AlibabaCampusDeviceOpenapiGetsimpledevicelistAPIResponse)
}

// ReleaseAlibabaCampusDeviceOpenapiGetsimpledevicelistAPIResponse 将 AlibabaCampusDeviceOpenapiGetsimpledevicelistAPIResponse 保存到 sync.Pool
func ReleaseAlibabaCampusDeviceOpenapiGetsimpledevicelistAPIResponse(v *AlibabaCampusDeviceOpenapiGetsimpledevicelistAPIResponse) {
	v.Reset()
	poolAlibabaCampusDeviceOpenapiGetsimpledevicelistAPIResponse.Put(v)
}
