package kclub

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaKclubKcQaSearchPageAPIResponse 知识云-知识检索(分页) API返回值
// alibaba.kclub.kc.qa.search.page
//
// 知识云-知识搜索服务
type AlibabaKclubKcQaSearchPageAPIResponse struct {
	model.CommonResponse
	AlibabaKclubKcQaSearchPageAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaKclubKcQaSearchPageAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaKclubKcQaSearchPageAPIResponseModel).Reset()
}

// AlibabaKclubKcQaSearchPageAPIResponseModel is 知识云-知识检索(分页) 成功返回结果
type AlibabaKclubKcQaSearchPageAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_kclub_kc_qa_search_page_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *AlibabaKclubKcQaSearchPageResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaKclubKcQaSearchPageAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaKclubKcQaSearchPageAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaKclubKcQaSearchPageAPIResponse)
	},
}

// GetAlibabaKclubKcQaSearchPageAPIResponse 从 sync.Pool 获取 AlibabaKclubKcQaSearchPageAPIResponse
func GetAlibabaKclubKcQaSearchPageAPIResponse() *AlibabaKclubKcQaSearchPageAPIResponse {
	return poolAlibabaKclubKcQaSearchPageAPIResponse.Get().(*AlibabaKclubKcQaSearchPageAPIResponse)
}

// ReleaseAlibabaKclubKcQaSearchPageAPIResponse 将 AlibabaKclubKcQaSearchPageAPIResponse 保存到 sync.Pool
func ReleaseAlibabaKclubKcQaSearchPageAPIResponse(v *AlibabaKclubKcQaSearchPageAPIResponse) {
	v.Reset()
	poolAlibabaKclubKcQaSearchPageAPIResponse.Put(v)
}
