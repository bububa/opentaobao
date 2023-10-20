package kclub

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaKclubKcQaSearchAPIResponse 知识云-知识检索 API返回值
// alibaba.kclub.kc.qa.search
//
// 知识云-知识搜索服务
type AlibabaKclubKcQaSearchAPIResponse struct {
	model.CommonResponse
	AlibabaKclubKcQaSearchAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaKclubKcQaSearchAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaKclubKcQaSearchAPIResponseModel).Reset()
}

// AlibabaKclubKcQaSearchAPIResponseModel is 知识云-知识检索 成功返回结果
type AlibabaKclubKcQaSearchAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_kclub_kc_qa_search_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 搜索结果
	Result *AlibabaKclubKcQaSearchResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaKclubKcQaSearchAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaKclubKcQaSearchAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaKclubKcQaSearchAPIResponse)
	},
}

// GetAlibabaKclubKcQaSearchAPIResponse 从 sync.Pool 获取 AlibabaKclubKcQaSearchAPIResponse
func GetAlibabaKclubKcQaSearchAPIResponse() *AlibabaKclubKcQaSearchAPIResponse {
	return poolAlibabaKclubKcQaSearchAPIResponse.Get().(*AlibabaKclubKcQaSearchAPIResponse)
}

// ReleaseAlibabaKclubKcQaSearchAPIResponse 将 AlibabaKclubKcQaSearchAPIResponse 保存到 sync.Pool
func ReleaseAlibabaKclubKcQaSearchAPIResponse(v *AlibabaKclubKcQaSearchAPIResponse) {
	v.Reset()
	poolAlibabaKclubKcQaSearchAPIResponse.Put(v)
}
