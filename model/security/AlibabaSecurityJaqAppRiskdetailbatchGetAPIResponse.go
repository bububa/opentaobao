package security

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabasecurityjaqappriskdetailbatchgetAPIResponse 应用风险详细信息批量查询接口 API返回值
// alibaba.security.jaq.app.riskdetailbatch.get
//
// 用户通过alibaba.security.jaq.app.risk.scanbatch接口提交应用进行风险批量扫描后，用此接口批量获取风险详细信息,包含漏洞列表、恶意代码列表、仿冒应用列表等信息
type AlibabasecurityjaqappriskdetailbatchgetAPIResponse struct {
	model.CommonResponse
	AlibabasecurityjaqappriskdetailbatchgetAPIResponseModel
}

// AlibabasecurityjaqappriskdetailbatchgetAPIResponseModel is 应用风险详细信息批量查询接口 成功返回结果
type AlibabasecurityjaqappriskdetailbatchgetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_security_jaq_app_riskdetailbatch_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 批量扫描风险详情
	Result *RiskDetailBatch `json:"result,omitempty" xml:"result,omitempty"`
}
