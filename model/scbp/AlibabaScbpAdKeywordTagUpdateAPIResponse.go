package scbp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpAdKeywordTagUpdateAPIResponse 修改关键词所属分组 API返回值
// alibaba.scbp.ad.keyword.tag.update
//
// 修改关键词所属分组
type AlibabaScbpAdKeywordTagUpdateAPIResponse struct {
	model.CommonResponse
	AlibabaScbpAdKeywordTagUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaScbpAdKeywordTagUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaScbpAdKeywordTagUpdateAPIResponseModel).Reset()
}

// AlibabaScbpAdKeywordTagUpdateAPIResponseModel is 修改关键词所属分组 成功返回结果
type AlibabaScbpAdKeywordTagUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_scbp_ad_keyword_tag_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 实际修改的关键词数
	Result int64 `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaScbpAdKeywordTagUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = 0
}

var poolAlibabaScbpAdKeywordTagUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaScbpAdKeywordTagUpdateAPIResponse)
	},
}

// GetAlibabaScbpAdKeywordTagUpdateAPIResponse 从 sync.Pool 获取 AlibabaScbpAdKeywordTagUpdateAPIResponse
func GetAlibabaScbpAdKeywordTagUpdateAPIResponse() *AlibabaScbpAdKeywordTagUpdateAPIResponse {
	return poolAlibabaScbpAdKeywordTagUpdateAPIResponse.Get().(*AlibabaScbpAdKeywordTagUpdateAPIResponse)
}

// ReleaseAlibabaScbpAdKeywordTagUpdateAPIResponse 将 AlibabaScbpAdKeywordTagUpdateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaScbpAdKeywordTagUpdateAPIResponse(v *AlibabaScbpAdKeywordTagUpdateAPIResponse) {
	v.Reset()
	poolAlibabaScbpAdKeywordTagUpdateAPIResponse.Put(v)
}
