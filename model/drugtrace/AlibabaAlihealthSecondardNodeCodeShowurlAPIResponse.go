package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthSecondardNodeCodeShowurlAPIResponse 查询码信息url API返回值
// alibaba.alihealth.secondard.node.code.showurl
//
// 二级节点查询码信息url
type AlibabaAlihealthSecondardNodeCodeShowurlAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthSecondardNodeCodeShowurlAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthSecondardNodeCodeShowurlAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthSecondardNodeCodeShowurlAPIResponseModel).Reset()
}

// AlibabaAlihealthSecondardNodeCodeShowurlAPIResponseModel is 查询码信息url 成功返回结果
type AlibabaAlihealthSecondardNodeCodeShowurlAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_secondard_node_code_showurl_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// model
	Model string `json:"model,omitempty" xml:"model,omitempty"`
	// msgInfo
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// msgCode
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 是否成功
	ResponseSuccess bool `json:"response_success,omitempty" xml:"response_success,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthSecondardNodeCodeShowurlAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Model = ""
	m.MsgInfo = ""
	m.MsgCode = ""
	m.ResponseSuccess = false
}

var poolAlibabaAlihealthSecondardNodeCodeShowurlAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthSecondardNodeCodeShowurlAPIResponse)
	},
}

// GetAlibabaAlihealthSecondardNodeCodeShowurlAPIResponse 从 sync.Pool 获取 AlibabaAlihealthSecondardNodeCodeShowurlAPIResponse
func GetAlibabaAlihealthSecondardNodeCodeShowurlAPIResponse() *AlibabaAlihealthSecondardNodeCodeShowurlAPIResponse {
	return poolAlibabaAlihealthSecondardNodeCodeShowurlAPIResponse.Get().(*AlibabaAlihealthSecondardNodeCodeShowurlAPIResponse)
}

// ReleaseAlibabaAlihealthSecondardNodeCodeShowurlAPIResponse 将 AlibabaAlihealthSecondardNodeCodeShowurlAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthSecondardNodeCodeShowurlAPIResponse(v *AlibabaAlihealthSecondardNodeCodeShowurlAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthSecondardNodeCodeShowurlAPIResponse.Put(v)
}
