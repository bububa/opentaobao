package scbp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpAdKeywordGetAPIResponse 外贸直通车查询关键词 API返回值
// alibaba.scbp.ad.keyword.get
//
// 外贸直通车查询关键词
type AlibabaScbpAdKeywordGetAPIResponse struct {
	model.CommonResponse
	AlibabaScbpAdKeywordGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaScbpAdKeywordGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaScbpAdKeywordGetAPIResponseModel).Reset()
}

// AlibabaScbpAdKeywordGetAPIResponseModel is 外贸直通车查询关键词 成功返回结果
type AlibabaScbpAdKeywordGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_scbp_ad_keyword_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 查询关键词列表
	KeywordList []KeywordResultDto `json:"keyword_list,omitempty" xml:"keyword_list>keyword_result_dto,omitempty"`
	// 总个数
	TotalNum int64 `json:"total_num,omitempty" xml:"total_num,omitempty"`
	// 总页数
	TotalPage int64 `json:"total_page,omitempty" xml:"total_page,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaScbpAdKeywordGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.KeywordList = m.KeywordList[:0]
	m.TotalNum = 0
	m.TotalPage = 0
}

var poolAlibabaScbpAdKeywordGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaScbpAdKeywordGetAPIResponse)
	},
}

// GetAlibabaScbpAdKeywordGetAPIResponse 从 sync.Pool 获取 AlibabaScbpAdKeywordGetAPIResponse
func GetAlibabaScbpAdKeywordGetAPIResponse() *AlibabaScbpAdKeywordGetAPIResponse {
	return poolAlibabaScbpAdKeywordGetAPIResponse.Get().(*AlibabaScbpAdKeywordGetAPIResponse)
}

// ReleaseAlibabaScbpAdKeywordGetAPIResponse 将 AlibabaScbpAdKeywordGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaScbpAdKeywordGetAPIResponse(v *AlibabaScbpAdKeywordGetAPIResponse) {
	v.Reset()
	poolAlibabaScbpAdKeywordGetAPIResponse.Put(v)
}
