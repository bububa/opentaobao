package campus

import (
	"encoding/xml"

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

// AlibabaCampusDeviceOpenapiGettemplatelistAPIResponseModel is 查询设备模板 成功返回结果
type AlibabaCampusDeviceOpenapiGettemplatelistAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_campus_device_openapi_gettemplatelist_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *PageResult `json:"result,omitempty" xml:"result,omitempty"`
}
