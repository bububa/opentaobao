package security

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSecurityJaqAppRiskdetailbatchGetAPIResponse 应用风险详细信息批量查询接口 API返回值
// alibaba.security.jaq.app.riskdetailbatch.get
//
// 用户通过alibaba.security.jaq.app.risk.scanbatch接口提交应用进行风险批量扫描后，用此接口批量获取风险详细信息,包含漏洞列表、恶意代码列表、仿冒应用列表等信息
type AlibabaSecurityJaqAppRiskdetailbatchGetAPIResponse struct {
	model.CommonResponse
	AlibabaSecurityJaqAppRiskdetailbatchGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaSecurityJaqAppRiskdetailbatchGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaSecurityJaqAppRiskdetailbatchGetAPIResponseModel).Reset()
}

// AlibabaSecurityJaqAppRiskdetailbatchGetAPIResponseModel is 应用风险详细信息批量查询接口 成功返回结果
type AlibabaSecurityJaqAppRiskdetailbatchGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_security_jaq_app_riskdetailbatch_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 批量扫描风险详情
	Result *RiskDetailBatch `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaSecurityJaqAppRiskdetailbatchGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaSecurityJaqAppRiskdetailbatchGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaSecurityJaqAppRiskdetailbatchGetAPIResponse)
	},
}

// GetAlibabaSecurityJaqAppRiskdetailbatchGetAPIResponse 从 sync.Pool 获取 AlibabaSecurityJaqAppRiskdetailbatchGetAPIResponse
func GetAlibabaSecurityJaqAppRiskdetailbatchGetAPIResponse() *AlibabaSecurityJaqAppRiskdetailbatchGetAPIResponse {
	return poolAlibabaSecurityJaqAppRiskdetailbatchGetAPIResponse.Get().(*AlibabaSecurityJaqAppRiskdetailbatchGetAPIResponse)
}

// ReleaseAlibabaSecurityJaqAppRiskdetailbatchGetAPIResponse 将 AlibabaSecurityJaqAppRiskdetailbatchGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaSecurityJaqAppRiskdetailbatchGetAPIResponse(v *AlibabaSecurityJaqAppRiskdetailbatchGetAPIResponse) {
	v.Reset()
	poolAlibabaSecurityJaqAppRiskdetailbatchGetAPIResponse.Put(v)
}
