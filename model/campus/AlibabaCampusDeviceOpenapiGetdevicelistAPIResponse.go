package campus

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusDeviceOpenapiGetdevicelistAPIResponse 多条件查询设备分组 API返回值
// alibaba.campus.device.openapi.getdevicelist
//
// 多条件查询设备分组
type AlibabaCampusDeviceOpenapiGetdevicelistAPIResponse struct {
	model.CommonResponse
	AlibabaCampusDeviceOpenapiGetdevicelistAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaCampusDeviceOpenapiGetdevicelistAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaCampusDeviceOpenapiGetdevicelistAPIResponseModel).Reset()
}

// AlibabaCampusDeviceOpenapiGetdevicelistAPIResponseModel is 多条件查询设备分组 成功返回结果
type AlibabaCampusDeviceOpenapiGetdevicelistAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_campus_device_openapi_getdevicelist_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回对象
	Result *PageResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaCampusDeviceOpenapiGetdevicelistAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaCampusDeviceOpenapiGetdevicelistAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaCampusDeviceOpenapiGetdevicelistAPIResponse)
	},
}

// GetAlibabaCampusDeviceOpenapiGetdevicelistAPIResponse 从 sync.Pool 获取 AlibabaCampusDeviceOpenapiGetdevicelistAPIResponse
func GetAlibabaCampusDeviceOpenapiGetdevicelistAPIResponse() *AlibabaCampusDeviceOpenapiGetdevicelistAPIResponse {
	return poolAlibabaCampusDeviceOpenapiGetdevicelistAPIResponse.Get().(*AlibabaCampusDeviceOpenapiGetdevicelistAPIResponse)
}

// ReleaseAlibabaCampusDeviceOpenapiGetdevicelistAPIResponse 将 AlibabaCampusDeviceOpenapiGetdevicelistAPIResponse 保存到 sync.Pool
func ReleaseAlibabaCampusDeviceOpenapiGetdevicelistAPIResponse(v *AlibabaCampusDeviceOpenapiGetdevicelistAPIResponse) {
	v.Reset()
	poolAlibabaCampusDeviceOpenapiGetdevicelistAPIResponse.Put(v)
}
