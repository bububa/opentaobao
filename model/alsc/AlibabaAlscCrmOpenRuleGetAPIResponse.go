package alsc

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscCrmOpenRuleGetAPIResponse 查询规则 API返回值
// alibaba.alsc.crm.open.rule.get
//
// 查询会员规则
type AlibabaAlscCrmOpenRuleGetAPIResponse struct {
	model.CommonResponse
	AlibabaAlscCrmOpenRuleGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlscCrmOpenRuleGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlscCrmOpenRuleGetAPIResponseModel).Reset()
}

// AlibabaAlscCrmOpenRuleGetAPIResponseModel is 查询规则 成功返回结果
type AlibabaAlscCrmOpenRuleGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alsc_crm_open_rule_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口结果
	Result *CommonResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlscCrmOpenRuleGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlscCrmOpenRuleGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlscCrmOpenRuleGetAPIResponse)
	},
}

// GetAlibabaAlscCrmOpenRuleGetAPIResponse 从 sync.Pool 获取 AlibabaAlscCrmOpenRuleGetAPIResponse
func GetAlibabaAlscCrmOpenRuleGetAPIResponse() *AlibabaAlscCrmOpenRuleGetAPIResponse {
	return poolAlibabaAlscCrmOpenRuleGetAPIResponse.Get().(*AlibabaAlscCrmOpenRuleGetAPIResponse)
}

// ReleaseAlibabaAlscCrmOpenRuleGetAPIResponse 将 AlibabaAlscCrmOpenRuleGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlscCrmOpenRuleGetAPIResponse(v *AlibabaAlscCrmOpenRuleGetAPIResponse) {
	v.Reset()
	poolAlibabaAlscCrmOpenRuleGetAPIResponse.Put(v)
}
