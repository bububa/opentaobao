package scbp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpReckeywordSearchAPIResponse 推荐词-词推词 API返回值
// alibaba.scbp.reckeyword.search
//
// 推荐词-词推词
type AlibabaScbpReckeywordSearchAPIResponse struct {
	model.CommonResponse
	AlibabaScbpReckeywordSearchAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaScbpReckeywordSearchAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaScbpReckeywordSearchAPIResponseModel).Reset()
}

// AlibabaScbpReckeywordSearchAPIResponseModel is 推荐词-词推词 成功返回结果
type AlibabaScbpReckeywordSearchAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_scbp_reckeyword_search_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 词推词结果列表
	ResultList []RecKeywordDto `json:"result_list,omitempty" xml:"result_list>rec_keyword_dto,omitempty"`
	// 总个数
	TotalNum int64 `json:"total_num,omitempty" xml:"total_num,omitempty"`
	// 总页数
	TotalPage int64 `json:"total_page,omitempty" xml:"total_page,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaScbpReckeywordSearchAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResultList = m.ResultList[:0]
	m.TotalNum = 0
	m.TotalPage = 0
}

var poolAlibabaScbpReckeywordSearchAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaScbpReckeywordSearchAPIResponse)
	},
}

// GetAlibabaScbpReckeywordSearchAPIResponse 从 sync.Pool 获取 AlibabaScbpReckeywordSearchAPIResponse
func GetAlibabaScbpReckeywordSearchAPIResponse() *AlibabaScbpReckeywordSearchAPIResponse {
	return poolAlibabaScbpReckeywordSearchAPIResponse.Get().(*AlibabaScbpReckeywordSearchAPIResponse)
}

// ReleaseAlibabaScbpReckeywordSearchAPIResponse 将 AlibabaScbpReckeywordSearchAPIResponse 保存到 sync.Pool
func ReleaseAlibabaScbpReckeywordSearchAPIResponse(v *AlibabaScbpReckeywordSearchAPIResponse) {
	v.Reset()
	poolAlibabaScbpReckeywordSearchAPIResponse.Put(v)
}
