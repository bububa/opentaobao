package alsc

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscCrmCardQueryTemplateAPIResponse 查询卡模板详情 API返回值
// alibaba.alsc.crm.card.query.template
//
// 查询卡模板详情
type AlibabaAlscCrmCardQueryTemplateAPIResponse struct {
	model.CommonResponse
	AlibabaAlscCrmCardQueryTemplateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlscCrmCardQueryTemplateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlscCrmCardQueryTemplateAPIResponseModel).Reset()
}

// AlibabaAlscCrmCardQueryTemplateAPIResponseModel is 查询卡模板详情 成功返回结果
type AlibabaAlscCrmCardQueryTemplateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alsc_crm_card_query_template_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口结果
	Result *CommonResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlscCrmCardQueryTemplateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlscCrmCardQueryTemplateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlscCrmCardQueryTemplateAPIResponse)
	},
}

// GetAlibabaAlscCrmCardQueryTemplateAPIResponse 从 sync.Pool 获取 AlibabaAlscCrmCardQueryTemplateAPIResponse
func GetAlibabaAlscCrmCardQueryTemplateAPIResponse() *AlibabaAlscCrmCardQueryTemplateAPIResponse {
	return poolAlibabaAlscCrmCardQueryTemplateAPIResponse.Get().(*AlibabaAlscCrmCardQueryTemplateAPIResponse)
}

// ReleaseAlibabaAlscCrmCardQueryTemplateAPIResponse 将 AlibabaAlscCrmCardQueryTemplateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlscCrmCardQueryTemplateAPIResponse(v *AlibabaAlscCrmCardQueryTemplateAPIResponse) {
	v.Reset()
	poolAlibabaAlscCrmCardQueryTemplateAPIResponse.Put(v)
}
