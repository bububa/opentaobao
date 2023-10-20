package alsc

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscCrmRuleQuerytaglistAPIResponse 查询标签列表 API返回值
// alibaba.alsc.crm.rule.querytaglist
//
// 查询标签列表
type AlibabaAlscCrmRuleQuerytaglistAPIResponse struct {
	model.CommonResponse
	AlibabaAlscCrmRuleQuerytaglistAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlscCrmRuleQuerytaglistAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlscCrmRuleQuerytaglistAPIResponseModel).Reset()
}

// AlibabaAlscCrmRuleQuerytaglistAPIResponseModel is 查询标签列表 成功返回结果
type AlibabaAlscCrmRuleQuerytaglistAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alsc_crm_rule_querytaglist_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 分页返回模型
	Result *CommonPageResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlscCrmRuleQuerytaglistAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlscCrmRuleQuerytaglistAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlscCrmRuleQuerytaglistAPIResponse)
	},
}

// GetAlibabaAlscCrmRuleQuerytaglistAPIResponse 从 sync.Pool 获取 AlibabaAlscCrmRuleQuerytaglistAPIResponse
func GetAlibabaAlscCrmRuleQuerytaglistAPIResponse() *AlibabaAlscCrmRuleQuerytaglistAPIResponse {
	return poolAlibabaAlscCrmRuleQuerytaglistAPIResponse.Get().(*AlibabaAlscCrmRuleQuerytaglistAPIResponse)
}

// ReleaseAlibabaAlscCrmRuleQuerytaglistAPIResponse 将 AlibabaAlscCrmRuleQuerytaglistAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlscCrmRuleQuerytaglistAPIResponse(v *AlibabaAlscCrmRuleQuerytaglistAPIResponse) {
	v.Reset()
	poolAlibabaAlscCrmRuleQuerytaglistAPIResponse.Put(v)
}
