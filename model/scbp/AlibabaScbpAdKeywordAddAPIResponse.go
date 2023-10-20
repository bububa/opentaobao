package scbp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpAdKeywordAddAPIResponse 外贸直通车加词 API返回值
// alibaba.scbp.ad.keyword.add
//
// 外贸直通车加词服务
type AlibabaScbpAdKeywordAddAPIResponse struct {
	model.CommonResponse
	AlibabaScbpAdKeywordAddAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaScbpAdKeywordAddAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaScbpAdKeywordAddAPIResponseModel).Reset()
}

// AlibabaScbpAdKeywordAddAPIResponseModel is 外贸直通车加词 成功返回结果
type AlibabaScbpAdKeywordAddAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_scbp_ad_keyword_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 请求加入的词
	Keyword string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// 加词失败的原因
	InvalidType string `json:"invalid_type,omitempty" xml:"invalid_type,omitempty"`
	// 系统中存在归一化重复的词
	RepeatKeyword string `json:"repeat_keyword,omitempty" xml:"repeat_keyword,omitempty"`
	// 该词是否加入成功
	IsAdded bool `json:"is_added,omitempty" xml:"is_added,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaScbpAdKeywordAddAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Keyword = ""
	m.InvalidType = ""
	m.RepeatKeyword = ""
	m.IsAdded = false
}

var poolAlibabaScbpAdKeywordAddAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaScbpAdKeywordAddAPIResponse)
	},
}

// GetAlibabaScbpAdKeywordAddAPIResponse 从 sync.Pool 获取 AlibabaScbpAdKeywordAddAPIResponse
func GetAlibabaScbpAdKeywordAddAPIResponse() *AlibabaScbpAdKeywordAddAPIResponse {
	return poolAlibabaScbpAdKeywordAddAPIResponse.Get().(*AlibabaScbpAdKeywordAddAPIResponse)
}

// ReleaseAlibabaScbpAdKeywordAddAPIResponse 将 AlibabaScbpAdKeywordAddAPIResponse 保存到 sync.Pool
func ReleaseAlibabaScbpAdKeywordAddAPIResponse(v *AlibabaScbpAdKeywordAddAPIResponse) {
	v.Reset()
	poolAlibabaScbpAdKeywordAddAPIResponse.Put(v)
}
