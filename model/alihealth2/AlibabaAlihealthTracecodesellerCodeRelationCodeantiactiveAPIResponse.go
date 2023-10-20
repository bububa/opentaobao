package alihealth2

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthTracecodesellerCodeRelationCodeantiactiveAPIResponse 解除码的关联关系 API返回值
// alibaba.alihealth.tracecodeseller.code.relation.codeantiactive
//
// 解除码的关联关系
type AlibabaAlihealthTracecodesellerCodeRelationCodeantiactiveAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthTracecodesellerCodeRelationCodeantiactiveAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthTracecodesellerCodeRelationCodeantiactiveAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthTracecodesellerCodeRelationCodeantiactiveAPIResponseModel).Reset()
}

// AlibabaAlihealthTracecodesellerCodeRelationCodeantiactiveAPIResponseModel is 解除码的关联关系 成功返回结果
type AlibabaAlihealthTracecodesellerCodeRelationCodeantiactiveAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_tracecodeseller_code_relation_codeantiactive_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 成功失败信息编码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 成功失败信息
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 成功失败标记
	Model bool `json:"model,omitempty" xml:"model,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthTracecodesellerCodeRelationCodeantiactiveAPIResponseModel) Reset() {
	m.RequestId = ""
	m.MsgCode = ""
	m.MsgInfo = ""
	m.Model = false
}

var poolAlibabaAlihealthTracecodesellerCodeRelationCodeantiactiveAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthTracecodesellerCodeRelationCodeantiactiveAPIResponse)
	},
}

// GetAlibabaAlihealthTracecodesellerCodeRelationCodeantiactiveAPIResponse 从 sync.Pool 获取 AlibabaAlihealthTracecodesellerCodeRelationCodeantiactiveAPIResponse
func GetAlibabaAlihealthTracecodesellerCodeRelationCodeantiactiveAPIResponse() *AlibabaAlihealthTracecodesellerCodeRelationCodeantiactiveAPIResponse {
	return poolAlibabaAlihealthTracecodesellerCodeRelationCodeantiactiveAPIResponse.Get().(*AlibabaAlihealthTracecodesellerCodeRelationCodeantiactiveAPIResponse)
}

// ReleaseAlibabaAlihealthTracecodesellerCodeRelationCodeantiactiveAPIResponse 将 AlibabaAlihealthTracecodesellerCodeRelationCodeantiactiveAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthTracecodesellerCodeRelationCodeantiactiveAPIResponse(v *AlibabaAlihealthTracecodesellerCodeRelationCodeantiactiveAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthTracecodesellerCodeRelationCodeantiactiveAPIResponse.Put(v)
}
