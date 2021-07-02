package campus

import (
	"encoding/xml"

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

// AlibabaCampusDeviceOpenapiOperatedeviceAPIResponseModel is 根据uuid操作设备 成功返回结果
type AlibabaCampusDeviceOpenapiOperatedeviceAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_campus_device_openapi_operatedevice_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回对象
	Result *PojoResult `json:"result,omitempty" xml:"result,omitempty"`
}
