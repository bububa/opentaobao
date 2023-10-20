package alihouse

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeAgreementPreshowAPIResponse 预览地址获取接口 API返回值
// alibaba.alihouse.newhome.agreement.preshow
//
// 预览地址获取接口
type AlibabaAlihouseNewhomeAgreementPreshowAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseNewhomeAgreementPreshowAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihouseNewhomeAgreementPreshowAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihouseNewhomeAgreementPreshowAPIResponseModel).Reset()
}

// AlibabaAlihouseNewhomeAgreementPreshowAPIResponseModel is 预览地址获取接口 成功返回结果
type AlibabaAlihouseNewhomeAgreementPreshowAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_agreement_preshow_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *AlibabaAlihouseNewhomeAgreementPreshowResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihouseNewhomeAgreementPreshowAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihouseNewhomeAgreementPreshowAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeAgreementPreshowAPIResponse)
	},
}

// GetAlibabaAlihouseNewhomeAgreementPreshowAPIResponse 从 sync.Pool 获取 AlibabaAlihouseNewhomeAgreementPreshowAPIResponse
func GetAlibabaAlihouseNewhomeAgreementPreshowAPIResponse() *AlibabaAlihouseNewhomeAgreementPreshowAPIResponse {
	return poolAlibabaAlihouseNewhomeAgreementPreshowAPIResponse.Get().(*AlibabaAlihouseNewhomeAgreementPreshowAPIResponse)
}

// ReleaseAlibabaAlihouseNewhomeAgreementPreshowAPIResponse 将 AlibabaAlihouseNewhomeAgreementPreshowAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihouseNewhomeAgreementPreshowAPIResponse(v *AlibabaAlihouseNewhomeAgreementPreshowAPIResponse) {
	v.Reset()
	poolAlibabaAlihouseNewhomeAgreementPreshowAPIResponse.Put(v)
}
