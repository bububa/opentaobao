package b2bcert

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAuthCertGetAPIResponse 获取证书数据 API返回值
// alibaba.auth.cert.get
//
// 获取证书数据
type AlibabaAuthCertGetAPIResponse struct {
	model.CommonResponse
	AlibabaAuthCertGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAuthCertGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAuthCertGetAPIResponseModel).Reset()
}

// AlibabaAuthCertGetAPIResponseModel is 获取证书数据 成功返回结果
type AlibabaAuthCertGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_auth_cert_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaAuthCertGetResponse `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAuthCertGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAuthCertGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAuthCertGetAPIResponse)
	},
}

// GetAlibabaAuthCertGetAPIResponse 从 sync.Pool 获取 AlibabaAuthCertGetAPIResponse
func GetAlibabaAuthCertGetAPIResponse() *AlibabaAuthCertGetAPIResponse {
	return poolAlibabaAuthCertGetAPIResponse.Get().(*AlibabaAuthCertGetAPIResponse)
}

// ReleaseAlibabaAuthCertGetAPIResponse 将 AlibabaAuthCertGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAuthCertGetAPIResponse(v *AlibabaAuthCertGetAPIResponse) {
	v.Reset()
	poolAlibabaAuthCertGetAPIResponse.Put(v)
}
