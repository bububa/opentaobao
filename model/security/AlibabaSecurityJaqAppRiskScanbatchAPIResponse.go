package security

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSecurityJaqAppRiskScanbatchAPIResponse 应用风险扫描批量提交接口 API返回值
// alibaba.security.jaq.app.risk.scanbatch
//
// 批量提交应用进行风险扫描(含漏洞扫描、恶意代码检测),扫描完成后可通过对应的查询接口查询扫描结果
type AlibabaSecurityJaqAppRiskScanbatchAPIResponse struct {
	model.CommonResponse
	AlibabaSecurityJaqAppRiskScanbatchAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaSecurityJaqAppRiskScanbatchAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaSecurityJaqAppRiskScanbatchAPIResponseModel).Reset()
}

// AlibabaSecurityJaqAppRiskScanbatchAPIResponseModel is 应用风险扫描批量提交接口 成功返回结果
type AlibabaSecurityJaqAppRiskScanbatchAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_security_jaq_app_risk_scanbatch_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 扫描任务信息
	Result *TaskInfo `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaSecurityJaqAppRiskScanbatchAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaSecurityJaqAppRiskScanbatchAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaSecurityJaqAppRiskScanbatchAPIResponse)
	},
}

// GetAlibabaSecurityJaqAppRiskScanbatchAPIResponse 从 sync.Pool 获取 AlibabaSecurityJaqAppRiskScanbatchAPIResponse
func GetAlibabaSecurityJaqAppRiskScanbatchAPIResponse() *AlibabaSecurityJaqAppRiskScanbatchAPIResponse {
	return poolAlibabaSecurityJaqAppRiskScanbatchAPIResponse.Get().(*AlibabaSecurityJaqAppRiskScanbatchAPIResponse)
}

// ReleaseAlibabaSecurityJaqAppRiskScanbatchAPIResponse 将 AlibabaSecurityJaqAppRiskScanbatchAPIResponse 保存到 sync.Pool
func ReleaseAlibabaSecurityJaqAppRiskScanbatchAPIResponse(v *AlibabaSecurityJaqAppRiskScanbatchAPIResponse) {
	v.Reset()
	poolAlibabaSecurityJaqAppRiskScanbatchAPIResponse.Put(v)
}
