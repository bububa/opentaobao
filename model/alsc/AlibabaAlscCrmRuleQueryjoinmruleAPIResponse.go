package alsc

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscCrmRuleQueryjoinmruleAPIResponse 查询品牌下的成为会员规则 API返回值
// alibaba.alsc.crm.rule.queryjoinmrule
//
// 查询品牌下的成为会员规则
type AlibabaAlscCrmRuleQueryjoinmruleAPIResponse struct {
	model.CommonResponse
	AlibabaAlscCrmRuleQueryjoinmruleAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlscCrmRuleQueryjoinmruleAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlscCrmRuleQueryjoinmruleAPIResponseModel).Reset()
}

// AlibabaAlscCrmRuleQueryjoinmruleAPIResponseModel is 查询品牌下的成为会员规则 成功返回结果
type AlibabaAlscCrmRuleQueryjoinmruleAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alsc_crm_rule_queryjoinmrule_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口结果
	Result *CommonResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlscCrmRuleQueryjoinmruleAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlscCrmRuleQueryjoinmruleAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlscCrmRuleQueryjoinmruleAPIResponse)
	},
}

// GetAlibabaAlscCrmRuleQueryjoinmruleAPIResponse 从 sync.Pool 获取 AlibabaAlscCrmRuleQueryjoinmruleAPIResponse
func GetAlibabaAlscCrmRuleQueryjoinmruleAPIResponse() *AlibabaAlscCrmRuleQueryjoinmruleAPIResponse {
	return poolAlibabaAlscCrmRuleQueryjoinmruleAPIResponse.Get().(*AlibabaAlscCrmRuleQueryjoinmruleAPIResponse)
}

// ReleaseAlibabaAlscCrmRuleQueryjoinmruleAPIResponse 将 AlibabaAlscCrmRuleQueryjoinmruleAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlscCrmRuleQueryjoinmruleAPIResponse(v *AlibabaAlscCrmRuleQueryjoinmruleAPIResponse) {
	v.Reset()
	poolAlibabaAlscCrmRuleQueryjoinmruleAPIResponse.Put(v)
}
