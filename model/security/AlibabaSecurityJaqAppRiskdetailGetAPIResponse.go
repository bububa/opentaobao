package security

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSecurityJaqAppRiskdetailGetAPIResponse 应用风险详细信息查询接口 API返回值
// alibaba.security.jaq.app.riskdetail.get
//
// 用户通过alibaba.security.jaq.app.risk.scan接口提交应用进行风险扫描后，用此接口获取风险详细信息,包含漏洞列表、恶意代码列表、仿冒应用列表等信息
type AlibabaSecurityJaqAppRiskdetailGetAPIResponse struct {
	model.CommonResponse
	AlibabaSecurityJaqAppRiskdetailGetAPIResponseModel
}

// AlibabaSecurityJaqAppRiskdetailGetAPIResponseModel is 应用风险详细信息查询接口 成功返回结果
type AlibabaSecurityJaqAppRiskdetailGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_security_jaq_app_riskdetail_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 风险详情
	Result *RiskDetail `json:"result,omitempty" xml:"result,omitempty"`
}
