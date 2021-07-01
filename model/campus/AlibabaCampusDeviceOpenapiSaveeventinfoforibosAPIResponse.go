package campus

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaCampusDeviceOpenapiSaveeventinfoforibosAPIResponse
saveeventinfoforibos API返回值
alibaba.campus.device.openapi.saveeventinfoforibos

IBos的事件信息上报与反馈的处理接口 */
type AlibabaCampusDeviceOpenapiSaveeventinfoforibosAPIResponse struct {
	model.CommonResponse
	AlibabaCampusDeviceOpenapiSaveeventinfoforibosAPIResponseModel
}

// AlibabaCampusDeviceOpenapiSaveeventinfoforibosAPIResponseModel is saveeventinfoforibos 成功返回结果
type AlibabaCampusDeviceOpenapiSaveeventinfoforibosAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_campus_device_openapi_saveeventinfoforibos_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// errorMsg
	RequestErrorMsg string `json:"request_error_msg,omitempty" xml:"request_error_msg,omitempty"`
	// success
	RequestSuccess bool `json:"request_success,omitempty" xml:"request_success,omitempty"`
}
