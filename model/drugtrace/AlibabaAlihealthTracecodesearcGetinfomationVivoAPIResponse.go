package drugtrace

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthtracecodesearcgetinfomationvivoAPIResponse 获取vivo banner API返回值
// alibaba.alihealth.tracecodesearc.getinfomation.vivo
//
// 获取vivo banner  url
type AlibabaalihealthtracecodesearcgetinfomationvivoAPIResponse struct {
	model.CommonResponse
	AlibabaalihealthtracecodesearcgetinfomationvivoAPIResponseModel
}

// AlibabaalihealthtracecodesearcgetinfomationvivoAPIResponseModel is 获取vivo banner 成功返回结果
type AlibabaalihealthtracecodesearcgetinfomationvivoAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_tracecodesearc_getinfomation_vivo_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// bannerURL
	Model string `json:"model,omitempty" xml:"model,omitempty"`
	// 操作说明
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 操作码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
}
