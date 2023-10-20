package alsc

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscCrmRuleQuerymdayeruleAPIResponse 查询品牌下的会员日规则 API返回值
// alibaba.alsc.crm.rule.querymdayerule
//
// 查询品牌下的会员日规则
type AlibabaAlscCrmRuleQuerymdayeruleAPIResponse struct {
	model.CommonResponse
	AlibabaAlscCrmRuleQuerymdayeruleAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlscCrmRuleQuerymdayeruleAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlscCrmRuleQuerymdayeruleAPIResponseModel).Reset()
}

// AlibabaAlscCrmRuleQuerymdayeruleAPIResponseModel is 查询品牌下的会员日规则 成功返回结果
type AlibabaAlscCrmRuleQuerymdayeruleAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alsc_crm_rule_querymdayerule_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口结果
	Result *CommonResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlscCrmRuleQuerymdayeruleAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlscCrmRuleQuerymdayeruleAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlscCrmRuleQuerymdayeruleAPIResponse)
	},
}

// GetAlibabaAlscCrmRuleQuerymdayeruleAPIResponse 从 sync.Pool 获取 AlibabaAlscCrmRuleQuerymdayeruleAPIResponse
func GetAlibabaAlscCrmRuleQuerymdayeruleAPIResponse() *AlibabaAlscCrmRuleQuerymdayeruleAPIResponse {
	return poolAlibabaAlscCrmRuleQuerymdayeruleAPIResponse.Get().(*AlibabaAlscCrmRuleQuerymdayeruleAPIResponse)
}

// ReleaseAlibabaAlscCrmRuleQuerymdayeruleAPIResponse 将 AlibabaAlscCrmRuleQuerymdayeruleAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlscCrmRuleQuerymdayeruleAPIResponse(v *AlibabaAlscCrmRuleQuerymdayeruleAPIResponse) {
	v.Reset()
	poolAlibabaAlscCrmRuleQuerymdayeruleAPIResponse.Put(v)
}
