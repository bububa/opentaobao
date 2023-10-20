package scbp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpAdKeywordGetKeywordCountByQueryAPIResponse 计划关键词数目 API返回值
// alibaba.scbp.ad.keyword.get.keyword.count.by.query
//
// 计划关键词数目
type AlibabaScbpAdKeywordGetKeywordCountByQueryAPIResponse struct {
	model.CommonResponse
	AlibabaScbpAdKeywordGetKeywordCountByQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaScbpAdKeywordGetKeywordCountByQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaScbpAdKeywordGetKeywordCountByQueryAPIResponseModel).Reset()
}

// AlibabaScbpAdKeywordGetKeywordCountByQueryAPIResponseModel is 计划关键词数目 成功返回结果
type AlibabaScbpAdKeywordGetKeywordCountByQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_scbp_ad_keyword_get_keyword_count_by_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	Result int64 `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaScbpAdKeywordGetKeywordCountByQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = 0
}

var poolAlibabaScbpAdKeywordGetKeywordCountByQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaScbpAdKeywordGetKeywordCountByQueryAPIResponse)
	},
}

// GetAlibabaScbpAdKeywordGetKeywordCountByQueryAPIResponse 从 sync.Pool 获取 AlibabaScbpAdKeywordGetKeywordCountByQueryAPIResponse
func GetAlibabaScbpAdKeywordGetKeywordCountByQueryAPIResponse() *AlibabaScbpAdKeywordGetKeywordCountByQueryAPIResponse {
	return poolAlibabaScbpAdKeywordGetKeywordCountByQueryAPIResponse.Get().(*AlibabaScbpAdKeywordGetKeywordCountByQueryAPIResponse)
}

// ReleaseAlibabaScbpAdKeywordGetKeywordCountByQueryAPIResponse 将 AlibabaScbpAdKeywordGetKeywordCountByQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaScbpAdKeywordGetKeywordCountByQueryAPIResponse(v *AlibabaScbpAdKeywordGetKeywordCountByQueryAPIResponse) {
	v.Reset()
	poolAlibabaScbpAdKeywordGetKeywordCountByQueryAPIResponse.Put(v)
}
