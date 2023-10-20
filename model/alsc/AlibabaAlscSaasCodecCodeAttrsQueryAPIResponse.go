package alsc

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscSaasCodecCodeAttrsQueryAPIResponse 码业务属性查询 API返回值
// alibaba.alsc.saas.codec.code.attrs.query
//
// 码业务属性查询
type AlibabaAlscSaasCodecCodeAttrsQueryAPIResponse struct {
	model.CommonResponse
	AlibabaAlscSaasCodecCodeAttrsQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlscSaasCodecCodeAttrsQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlscSaasCodecCodeAttrsQueryAPIResponseModel).Reset()
}

// AlibabaAlscSaasCodecCodeAttrsQueryAPIResponseModel is 码业务属性查询 成功返回结果
type AlibabaAlscSaasCodecCodeAttrsQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alsc_saas_codec_code_attrs_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlscSaasCodecCodeAttrsQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlscSaasCodecCodeAttrsQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlscSaasCodecCodeAttrsQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlscSaasCodecCodeAttrsQueryAPIResponse)
	},
}

// GetAlibabaAlscSaasCodecCodeAttrsQueryAPIResponse 从 sync.Pool 获取 AlibabaAlscSaasCodecCodeAttrsQueryAPIResponse
func GetAlibabaAlscSaasCodecCodeAttrsQueryAPIResponse() *AlibabaAlscSaasCodecCodeAttrsQueryAPIResponse {
	return poolAlibabaAlscSaasCodecCodeAttrsQueryAPIResponse.Get().(*AlibabaAlscSaasCodecCodeAttrsQueryAPIResponse)
}

// ReleaseAlibabaAlscSaasCodecCodeAttrsQueryAPIResponse 将 AlibabaAlscSaasCodecCodeAttrsQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlscSaasCodecCodeAttrsQueryAPIResponse(v *AlibabaAlscSaasCodecCodeAttrsQueryAPIResponse) {
	v.Reset()
	poolAlibabaAlscSaasCodecCodeAttrsQueryAPIResponse.Put(v)
}
