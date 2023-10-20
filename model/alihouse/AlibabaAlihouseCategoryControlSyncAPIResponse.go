package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousecategorycontrolsyncAPIResponse 类目权限上翻 API返回值
// alibaba.alihouse.category.control.sync
//
// 类目权限上翻
type AlibabaalihousecategorycontrolsyncAPIResponse struct {
	model.CommonResponse
	AlibabaalihousecategorycontrolsyncAPIResponseModel
}

// AlibabaalihousecategorycontrolsyncAPIResponseModel is 类目权限上翻 成功返回结果
type AlibabaalihousecategorycontrolsyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_category_control_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// list
	Data []CategoryControlResultDto `json:"data,omitempty" xml:"data>category_control_result_dto,omitempty"`
	// code
	ReturnCode string `json:"return_code,omitempty" xml:"return_code,omitempty"`
	// message
	ReturnMessage string `json:"return_message,omitempty" xml:"return_message,omitempty"`
	// success
	ReturnSuccess bool `json:"return_success,omitempty" xml:"return_success,omitempty"`
}
