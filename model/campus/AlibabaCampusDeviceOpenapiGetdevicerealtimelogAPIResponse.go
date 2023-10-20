package campus

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusDeviceOpenapiGetdevicerealtimelogAPIResponse 根据设备uuid获取设备采集信息 API返回值
// alibaba.campus.device.openapi.getdevicerealtimelog
//
// 根据设备uuid获取设备采集信息
type AlibabaCampusDeviceOpenapiGetdevicerealtimelogAPIResponse struct {
	model.CommonResponse
	AlibabaCampusDeviceOpenapiGetdevicerealtimelogAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaCampusDeviceOpenapiGetdevicerealtimelogAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaCampusDeviceOpenapiGetdevicerealtimelogAPIResponseModel).Reset()
}

// AlibabaCampusDeviceOpenapiGetdevicerealtimelogAPIResponseModel is 根据设备uuid获取设备采集信息 成功返回结果
type AlibabaCampusDeviceOpenapiGetdevicerealtimelogAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_campus_device_openapi_getdevicerealtimelog_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回对象
	Result *PojoResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaCampusDeviceOpenapiGetdevicerealtimelogAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaCampusDeviceOpenapiGetdevicerealtimelogAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaCampusDeviceOpenapiGetdevicerealtimelogAPIResponse)
	},
}

// GetAlibabaCampusDeviceOpenapiGetdevicerealtimelogAPIResponse 从 sync.Pool 获取 AlibabaCampusDeviceOpenapiGetdevicerealtimelogAPIResponse
func GetAlibabaCampusDeviceOpenapiGetdevicerealtimelogAPIResponse() *AlibabaCampusDeviceOpenapiGetdevicerealtimelogAPIResponse {
	return poolAlibabaCampusDeviceOpenapiGetdevicerealtimelogAPIResponse.Get().(*AlibabaCampusDeviceOpenapiGetdevicerealtimelogAPIResponse)
}

// ReleaseAlibabaCampusDeviceOpenapiGetdevicerealtimelogAPIResponse 将 AlibabaCampusDeviceOpenapiGetdevicerealtimelogAPIResponse 保存到 sync.Pool
func ReleaseAlibabaCampusDeviceOpenapiGetdevicerealtimelogAPIResponse(v *AlibabaCampusDeviceOpenapiGetdevicerealtimelogAPIResponse) {
	v.Reset()
	poolAlibabaCampusDeviceOpenapiGetdevicerealtimelogAPIResponse.Put(v)
}
