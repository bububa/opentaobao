package campus

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusDeviceOpenapiOperatedeviceAPIResponse 根据uuid操作设备 API返回值
// alibaba.campus.device.openapi.operatedevice
//
// 根据uuid操作设备
type AlibabaCampusDeviceOpenapiOperatedeviceAPIResponse struct {
	model.CommonResponse
	AlibabaCampusDeviceOpenapiOperatedeviceAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaCampusDeviceOpenapiOperatedeviceAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaCampusDeviceOpenapiOperatedeviceAPIResponseModel).Reset()
}

// AlibabaCampusDeviceOpenapiOperatedeviceAPIResponseModel is 根据uuid操作设备 成功返回结果
type AlibabaCampusDeviceOpenapiOperatedeviceAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_campus_device_openapi_operatedevice_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回对象
	Result *PojoResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaCampusDeviceOpenapiOperatedeviceAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaCampusDeviceOpenapiOperatedeviceAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaCampusDeviceOpenapiOperatedeviceAPIResponse)
	},
}

// GetAlibabaCampusDeviceOpenapiOperatedeviceAPIResponse 从 sync.Pool 获取 AlibabaCampusDeviceOpenapiOperatedeviceAPIResponse
func GetAlibabaCampusDeviceOpenapiOperatedeviceAPIResponse() *AlibabaCampusDeviceOpenapiOperatedeviceAPIResponse {
	return poolAlibabaCampusDeviceOpenapiOperatedeviceAPIResponse.Get().(*AlibabaCampusDeviceOpenapiOperatedeviceAPIResponse)
}

// ReleaseAlibabaCampusDeviceOpenapiOperatedeviceAPIResponse 将 AlibabaCampusDeviceOpenapiOperatedeviceAPIResponse 保存到 sync.Pool
func ReleaseAlibabaCampusDeviceOpenapiOperatedeviceAPIResponse(v *AlibabaCampusDeviceOpenapiOperatedeviceAPIResponse) {
	v.Reset()
	poolAlibabaCampusDeviceOpenapiOperatedeviceAPIResponse.Put(v)
}
