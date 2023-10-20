package security

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSecurityJaqAppRisksummaryGetAPIResponse 应用风险概要信息查询接口 API返回值
// alibaba.security.jaq.app.risksummary.get
//
// 用户通过alibaba.security.jaq.app.risk.scan接口提交应用进行风险扫描后，用此接口获取风险概要信息，本接口不返回风险详细信息
type AlibabaSecurityJaqAppRisksummaryGetAPIResponse struct {
	model.CommonResponse
	AlibabaSecurityJaqAppRisksummaryGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaSecurityJaqAppRisksummaryGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaSecurityJaqAppRisksummaryGetAPIResponseModel).Reset()
}

// AlibabaSecurityJaqAppRisksummaryGetAPIResponseModel is 应用风险概要信息查询接口 成功返回结果
type AlibabaSecurityJaqAppRisksummaryGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_security_jaq_app_risksummary_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 应用扫描概要信息
	Result *RiskSummary `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaSecurityJaqAppRisksummaryGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaSecurityJaqAppRisksummaryGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaSecurityJaqAppRisksummaryGetAPIResponse)
	},
}

// GetAlibabaSecurityJaqAppRisksummaryGetAPIResponse 从 sync.Pool 获取 AlibabaSecurityJaqAppRisksummaryGetAPIResponse
func GetAlibabaSecurityJaqAppRisksummaryGetAPIResponse() *AlibabaSecurityJaqAppRisksummaryGetAPIResponse {
	return poolAlibabaSecurityJaqAppRisksummaryGetAPIResponse.Get().(*AlibabaSecurityJaqAppRisksummaryGetAPIResponse)
}

// ReleaseAlibabaSecurityJaqAppRisksummaryGetAPIResponse 将 AlibabaSecurityJaqAppRisksummaryGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaSecurityJaqAppRisksummaryGetAPIResponse(v *AlibabaSecurityJaqAppRisksummaryGetAPIResponse) {
	v.Reset()
	poolAlibabaSecurityJaqAppRisksummaryGetAPIResponse.Put(v)
}
