package campus

import (
	"encoding/xml"

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

// AlibabaCampusDeviceOpenapiGetuniquedeviceAPIResponseModel is 根据设备uuid获取设备信息 成功返回结果
type AlibabaCampusDeviceOpenapiGetuniquedeviceAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_campus_device_openapi_getuniquedevice_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回对象
	Result *PojoResult `json:"result,omitempty" xml:"result,omitempty"`
}
