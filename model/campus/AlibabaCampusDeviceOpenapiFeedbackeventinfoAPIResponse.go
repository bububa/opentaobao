package campus

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaCampusDeviceOpenapiFeedbackeventinfoAPIResponse
IVS事件处理反馈接口 API返回值
alibaba.campus.device.openapi.feedbackeventinfo

提供给第三方ISV的的事件信息处理反馈的接口 */
type AlibabaCampusDeviceOpenapiFeedbackeventinfoAPIResponse struct {
	model.CommonResponse
	AlibabaCampusDeviceOpenapiFeedbackeventinfoAPIResponseModel
}

// AlibabaCampusDeviceOpenapiFeedbackeventinfoAPIResponseModel is IVS事件处理反馈接口 成功返回结果
type AlibabaCampusDeviceOpenapiFeedbackeventinfoAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_campus_device_openapi_feedbackeventinfo_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *BaseResult `json:"result,omitempty" xml:"result,omitempty"`
}
