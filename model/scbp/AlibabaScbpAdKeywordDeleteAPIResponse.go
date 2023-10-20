package scbp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpAdKeywordDeleteAPIResponse 外贸直通车删除关键词 API返回值
// alibaba.scbp.ad.keyword.delete
//
// 外贸直通车删除关键词
type AlibabaScbpAdKeywordDeleteAPIResponse struct {
	model.CommonResponse
	AlibabaScbpAdKeywordDeleteAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaScbpAdKeywordDeleteAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaScbpAdKeywordDeleteAPIResponseModel).Reset()
}

// AlibabaScbpAdKeywordDeleteAPIResponseModel is 外贸直通车删除关键词 成功返回结果
type AlibabaScbpAdKeywordDeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_scbp_ad_keyword_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 删除关键词是否成功
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaScbpAdKeywordDeleteAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = false
}

var poolAlibabaScbpAdKeywordDeleteAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaScbpAdKeywordDeleteAPIResponse)
	},
}

// GetAlibabaScbpAdKeywordDeleteAPIResponse 从 sync.Pool 获取 AlibabaScbpAdKeywordDeleteAPIResponse
func GetAlibabaScbpAdKeywordDeleteAPIResponse() *AlibabaScbpAdKeywordDeleteAPIResponse {
	return poolAlibabaScbpAdKeywordDeleteAPIResponse.Get().(*AlibabaScbpAdKeywordDeleteAPIResponse)
}

// ReleaseAlibabaScbpAdKeywordDeleteAPIResponse 将 AlibabaScbpAdKeywordDeleteAPIResponse 保存到 sync.Pool
func ReleaseAlibabaScbpAdKeywordDeleteAPIResponse(v *AlibabaScbpAdKeywordDeleteAPIResponse) {
	v.Reset()
	poolAlibabaScbpAdKeywordDeleteAPIResponse.Put(v)
}
