package tmallservice

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallservicecodeconsumeAPIResponse 天猫服务平台服务核销 API返回值
// tmall.service.code.consume
//
// 天猫服务平台－服务核销
type TmallservicecodeconsumeAPIResponse struct {
	model.CommonResponse
	TmallservicecodeconsumeAPIResponseModel
}

// TmallservicecodeconsumeAPIResponseModel is 天猫服务平台服务核销 成功返回结果
type TmallservicecodeconsumeAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_service_code_consume_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 服务工单ID
	Result int64 `json:"result,omitempty" xml:"result,omitempty"`
}
