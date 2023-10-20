package alihealth2

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthtracecodeplatformcodeactiveAPIResponse 正大鸡蛋激活追溯码 API返回值
// alibaba.alihealth.tracecodeplatform.code.active
//
// 用于正大鸡蛋激活追溯码
type AlibabaalihealthtracecodeplatformcodeactiveAPIResponse struct {
	model.CommonResponse
	AlibabaalihealthtracecodeplatformcodeactiveAPIResponseModel
}

// AlibabaalihealthtracecodeplatformcodeactiveAPIResponseModel is 正大鸡蛋激活追溯码 成功返回结果
type AlibabaalihealthtracecodeplatformcodeactiveAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_tracecodeplatform_code_active_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *AlibabaalihealthtracecodeplatformcodeactiveResult `json:"result,omitempty" xml:"result,omitempty"`
}
