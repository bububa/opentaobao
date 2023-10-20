package alsc

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscCrmRuleQuerydishruleAPIResponse 查询品牌下的入会菜品规则 API返回值
// alibaba.alsc.crm.rule.querydishrule
//
// 查询品牌下的入会菜品规则
type AlibabaAlscCrmRuleQuerydishruleAPIResponse struct {
	model.CommonResponse
	AlibabaAlscCrmRuleQuerydishruleAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlscCrmRuleQuerydishruleAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlscCrmRuleQuerydishruleAPIResponseModel).Reset()
}

// AlibabaAlscCrmRuleQuerydishruleAPIResponseModel is 查询品牌下的入会菜品规则 成功返回结果
type AlibabaAlscCrmRuleQuerydishruleAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alsc_crm_rule_querydishrule_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 分页返回模型
	Result *CommonPageResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlscCrmRuleQuerydishruleAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlscCrmRuleQuerydishruleAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlscCrmRuleQuerydishruleAPIResponse)
	},
}

// GetAlibabaAlscCrmRuleQuerydishruleAPIResponse 从 sync.Pool 获取 AlibabaAlscCrmRuleQuerydishruleAPIResponse
func GetAlibabaAlscCrmRuleQuerydishruleAPIResponse() *AlibabaAlscCrmRuleQuerydishruleAPIResponse {
	return poolAlibabaAlscCrmRuleQuerydishruleAPIResponse.Get().(*AlibabaAlscCrmRuleQuerydishruleAPIResponse)
}

// ReleaseAlibabaAlscCrmRuleQuerydishruleAPIResponse 将 AlibabaAlscCrmRuleQuerydishruleAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlscCrmRuleQuerydishruleAPIResponse(v *AlibabaAlscCrmRuleQuerydishruleAPIResponse) {
	v.Reset()
	poolAlibabaAlscCrmRuleQuerydishruleAPIResponse.Put(v)
}
