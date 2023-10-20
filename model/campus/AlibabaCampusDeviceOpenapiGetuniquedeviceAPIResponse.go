package campus

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusDeviceOpenapiGetuniquedeviceAPIResponse 根据设备uuid获取设备信息 API返回值
// alibaba.campus.device.openapi.getuniquedevice
//
// 根据设备uuid获取设备信息
type AlibabaCampusDeviceOpenapiGetuniquedeviceAPIResponse struct {
	model.CommonResponse
	AlibabaCampusDeviceOpenapiGetuniquedeviceAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaCampusDeviceOpenapiGetuniquedeviceAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaCampusDeviceOpenapiGetuniquedeviceAPIResponseModel).Reset()
}

// AlibabaCampusDeviceOpenapiGetuniquedeviceAPIResponseModel is 根据设备uuid获取设备信息 成功返回结果
type AlibabaCampusDeviceOpenapiGetuniquedeviceAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_campus_device_openapi_getuniquedevice_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回对象
	Result *PojoResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaCampusDeviceOpenapiGetuniquedeviceAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaCampusDeviceOpenapiGetuniquedeviceAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaCampusDeviceOpenapiGetuniquedeviceAPIResponse)
	},
}

// GetAlibabaCampusDeviceOpenapiGetuniquedeviceAPIResponse 从 sync.Pool 获取 AlibabaCampusDeviceOpenapiGetuniquedeviceAPIResponse
func GetAlibabaCampusDeviceOpenapiGetuniquedeviceAPIResponse() *AlibabaCampusDeviceOpenapiGetuniquedeviceAPIResponse {
	return poolAlibabaCampusDeviceOpenapiGetuniquedeviceAPIResponse.Get().(*AlibabaCampusDeviceOpenapiGetuniquedeviceAPIResponse)
}

// ReleaseAlibabaCampusDeviceOpenapiGetuniquedeviceAPIResponse 将 AlibabaCampusDeviceOpenapiGetuniquedeviceAPIResponse 保存到 sync.Pool
func ReleaseAlibabaCampusDeviceOpenapiGetuniquedeviceAPIResponse(v *AlibabaCampusDeviceOpenapiGetuniquedeviceAPIResponse) {
	v.Reset()
	poolAlibabaCampusDeviceOpenapiGetuniquedeviceAPIResponse.Put(v)
}
