package campus

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusDeviceOpenapiGetdevicerealtimedataAPIResponse 获取指定设备下指定参数的实时值 API返回值
// alibaba.campus.device.openapi.getdevicerealtimedata
//
// 获取指定设备下指定参数的实时值
type AlibabaCampusDeviceOpenapiGetdevicerealtimedataAPIResponse struct {
	model.CommonResponse
	AlibabaCampusDeviceOpenapiGetdevicerealtimedataAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaCampusDeviceOpenapiGetdevicerealtimedataAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaCampusDeviceOpenapiGetdevicerealtimedataAPIResponseModel).Reset()
}

// AlibabaCampusDeviceOpenapiGetdevicerealtimedataAPIResponseModel is 获取指定设备下指定参数的实时值 成功返回结果
type AlibabaCampusDeviceOpenapiGetdevicerealtimedataAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_campus_device_openapi_getdevicerealtimedata_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回查询结果
	Result *PojoResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaCampusDeviceOpenapiGetdevicerealtimedataAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaCampusDeviceOpenapiGetdevicerealtimedataAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaCampusDeviceOpenapiGetdevicerealtimedataAPIResponse)
	},
}

// GetAlibabaCampusDeviceOpenapiGetdevicerealtimedataAPIResponse 从 sync.Pool 获取 AlibabaCampusDeviceOpenapiGetdevicerealtimedataAPIResponse
func GetAlibabaCampusDeviceOpenapiGetdevicerealtimedataAPIResponse() *AlibabaCampusDeviceOpenapiGetdevicerealtimedataAPIResponse {
	return poolAlibabaCampusDeviceOpenapiGetdevicerealtimedataAPIResponse.Get().(*AlibabaCampusDeviceOpenapiGetdevicerealtimedataAPIResponse)
}

// ReleaseAlibabaCampusDeviceOpenapiGetdevicerealtimedataAPIResponse 将 AlibabaCampusDeviceOpenapiGetdevicerealtimedataAPIResponse 保存到 sync.Pool
func ReleaseAlibabaCampusDeviceOpenapiGetdevicerealtimedataAPIResponse(v *AlibabaCampusDeviceOpenapiGetdevicerealtimedataAPIResponse) {
	v.Reset()
	poolAlibabaCampusDeviceOpenapiGetdevicerealtimedataAPIResponse.Put(v)
}
