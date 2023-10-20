package alsc

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscCrmRuleQueryoptplanAPIResponse 查询运营计划 API返回值
// alibaba.alsc.crm.rule.queryoptplan
//
// 查询运营计划
type AlibabaAlscCrmRuleQueryoptplanAPIResponse struct {
	model.CommonResponse
	AlibabaAlscCrmRuleQueryoptplanAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlscCrmRuleQueryoptplanAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlscCrmRuleQueryoptplanAPIResponseModel).Reset()
}

// AlibabaAlscCrmRuleQueryoptplanAPIResponseModel is 查询运营计划 成功返回结果
type AlibabaAlscCrmRuleQueryoptplanAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alsc_crm_rule_queryoptplan_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 请求参数
	Result *CommonPageResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlscCrmRuleQueryoptplanAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlscCrmRuleQueryoptplanAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlscCrmRuleQueryoptplanAPIResponse)
	},
}

// GetAlibabaAlscCrmRuleQueryoptplanAPIResponse 从 sync.Pool 获取 AlibabaAlscCrmRuleQueryoptplanAPIResponse
func GetAlibabaAlscCrmRuleQueryoptplanAPIResponse() *AlibabaAlscCrmRuleQueryoptplanAPIResponse {
	return poolAlibabaAlscCrmRuleQueryoptplanAPIResponse.Get().(*AlibabaAlscCrmRuleQueryoptplanAPIResponse)
}

// ReleaseAlibabaAlscCrmRuleQueryoptplanAPIResponse 将 AlibabaAlscCrmRuleQueryoptplanAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlscCrmRuleQueryoptplanAPIResponse(v *AlibabaAlscCrmRuleQueryoptplanAPIResponse) {
	v.Reset()
	poolAlibabaAlscCrmRuleQueryoptplanAPIResponse.Put(v)
}
