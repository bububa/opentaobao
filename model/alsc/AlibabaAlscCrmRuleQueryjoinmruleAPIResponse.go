package alsc

import (
	"encoding/xml"

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

// AlibabaAlscCrmRuleQueryjoinmruleAPIResponseModel is 查询品牌下的成为会员规则 成功返回结果
type AlibabaAlscCrmRuleQueryjoinmruleAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alsc_crm_rule_queryjoinmrule_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口结果
	Result *CommonResult `json:"result,omitempty" xml:"result,omitempty"`
}
