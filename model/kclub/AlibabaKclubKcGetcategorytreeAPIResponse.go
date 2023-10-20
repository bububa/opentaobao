package kclub

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaKclubKcGetcategorytreeAPIResponse 知识云-查询租户下类目树 API返回值
// alibaba.kclub.kc.getcategorytree
//
// 知识云-查询租户下类目树。通过租户id、类型(外部类目、帮助中心类目、内部类目)。
type AlibabaKclubKcGetcategorytreeAPIResponse struct {
	model.CommonResponse
	AlibabaKclubKcGetcategorytreeAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaKclubKcGetcategorytreeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaKclubKcGetcategorytreeAPIResponseModel).Reset()
}

// AlibabaKclubKcGetcategorytreeAPIResponseModel is 知识云-查询租户下类目树 成功返回结果
type AlibabaKclubKcGetcategorytreeAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_kclub_kc_getcategorytree_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaKclubKcGetcategorytreeResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaKclubKcGetcategorytreeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaKclubKcGetcategorytreeAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaKclubKcGetcategorytreeAPIResponse)
	},
}

// GetAlibabaKclubKcGetcategorytreeAPIResponse 从 sync.Pool 获取 AlibabaKclubKcGetcategorytreeAPIResponse
func GetAlibabaKclubKcGetcategorytreeAPIResponse() *AlibabaKclubKcGetcategorytreeAPIResponse {
	return poolAlibabaKclubKcGetcategorytreeAPIResponse.Get().(*AlibabaKclubKcGetcategorytreeAPIResponse)
}

// ReleaseAlibabaKclubKcGetcategorytreeAPIResponse 将 AlibabaKclubKcGetcategorytreeAPIResponse 保存到 sync.Pool
func ReleaseAlibabaKclubKcGetcategorytreeAPIResponse(v *AlibabaKclubKcGetcategorytreeAPIResponse) {
	v.Reset()
	poolAlibabaKclubKcGetcategorytreeAPIResponse.Put(v)
}
