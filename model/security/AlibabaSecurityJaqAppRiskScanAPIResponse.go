package security

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSecurityJaqAppRiskScanAPIResponse 应用风险扫描提交接口 API返回值
// alibaba.security.jaq.app.risk.scan
//
// 提交应用进行风险扫描(含漏洞扫描、恶意代码检测、仿冒监测),扫描完成后可通过对应的查询接口查询扫描结果
type AlibabaSecurityJaqAppRiskScanAPIResponse struct {
	model.CommonResponse
	AlibabaSecurityJaqAppRiskScanAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaSecurityJaqAppRiskScanAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaSecurityJaqAppRiskScanAPIResponseModel).Reset()
}

// AlibabaSecurityJaqAppRiskScanAPIResponseModel is 应用风险扫描提交接口 成功返回结果
type AlibabaSecurityJaqAppRiskScanAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_security_jaq_app_risk_scan_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 扫描任务信息
	Result *TaskInfo `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaSecurityJaqAppRiskScanAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaSecurityJaqAppRiskScanAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaSecurityJaqAppRiskScanAPIResponse)
	},
}

// GetAlibabaSecurityJaqAppRiskScanAPIResponse 从 sync.Pool 获取 AlibabaSecurityJaqAppRiskScanAPIResponse
func GetAlibabaSecurityJaqAppRiskScanAPIResponse() *AlibabaSecurityJaqAppRiskScanAPIResponse {
	return poolAlibabaSecurityJaqAppRiskScanAPIResponse.Get().(*AlibabaSecurityJaqAppRiskScanAPIResponse)
}

// ReleaseAlibabaSecurityJaqAppRiskScanAPIResponse 将 AlibabaSecurityJaqAppRiskScanAPIResponse 保存到 sync.Pool
func ReleaseAlibabaSecurityJaqAppRiskScanAPIResponse(v *AlibabaSecurityJaqAppRiskScanAPIResponse) {
	v.Reset()
	poolAlibabaSecurityJaqAppRiskScanAPIResponse.Put(v)
}
