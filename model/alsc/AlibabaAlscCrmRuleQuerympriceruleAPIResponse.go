package alsc

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscCrmRuleQuerympriceruleAPIResponse 查询品牌下的会员价规则 API返回值
// alibaba.alsc.crm.rule.querympricerule
//
// 查询品牌下的会员价规则
type AlibabaAlscCrmRuleQuerympriceruleAPIResponse struct {
	model.CommonResponse
	AlibabaAlscCrmRuleQuerympriceruleAPIResponseModel
}

// AlibabaAlscCrmRuleQuerympriceruleAPIResponseModel is 查询品牌下的会员价规则 成功返回结果
type AlibabaAlscCrmRuleQuerympriceruleAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alsc_crm_rule_querympricerule_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口结果
	Result *CommonResult `json:"result,omitempty" xml:"result,omitempty"`
}
