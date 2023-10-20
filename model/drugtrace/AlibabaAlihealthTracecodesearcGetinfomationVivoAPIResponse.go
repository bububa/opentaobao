package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthTracecodesearcGetinfomationVivoAPIResponse 获取vivo banner API返回值
// alibaba.alihealth.tracecodesearc.getinfomation.vivo
//
// 获取vivo banner  url
type AlibabaAlihealthTracecodesearcGetinfomationVivoAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthTracecodesearcGetinfomationVivoAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthTracecodesearcGetinfomationVivoAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthTracecodesearcGetinfomationVivoAPIResponseModel).Reset()
}

// AlibabaAlihealthTracecodesearcGetinfomationVivoAPIResponseModel is 获取vivo banner 成功返回结果
type AlibabaAlihealthTracecodesearcGetinfomationVivoAPIResponseModel struct {
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

// Reset 清空结构体
func (m *AlibabaAlihealthTracecodesearcGetinfomationVivoAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Model = ""
	m.MsgInfo = ""
	m.MsgCode = ""
}

var poolAlibabaAlihealthTracecodesearcGetinfomationVivoAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthTracecodesearcGetinfomationVivoAPIResponse)
	},
}

// GetAlibabaAlihealthTracecodesearcGetinfomationVivoAPIResponse 从 sync.Pool 获取 AlibabaAlihealthTracecodesearcGetinfomationVivoAPIResponse
func GetAlibabaAlihealthTracecodesearcGetinfomationVivoAPIResponse() *AlibabaAlihealthTracecodesearcGetinfomationVivoAPIResponse {
	return poolAlibabaAlihealthTracecodesearcGetinfomationVivoAPIResponse.Get().(*AlibabaAlihealthTracecodesearcGetinfomationVivoAPIResponse)
}

// ReleaseAlibabaAlihealthTracecodesearcGetinfomationVivoAPIResponse 将 AlibabaAlihealthTracecodesearcGetinfomationVivoAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthTracecodesearcGetinfomationVivoAPIResponse(v *AlibabaAlihealthTracecodesearcGetinfomationVivoAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthTracecodesearcGetinfomationVivoAPIResponse.Put(v)
}
