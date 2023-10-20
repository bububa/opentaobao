package scbp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpAdKeywordRankGetAPIResponse 获取外贸直通车关键词预估排名 API返回值
// alibaba.scbp.ad.keyword.rank.get
//
// 获取外贸直通车关键词预估排名
type AlibabaScbpAdKeywordRankGetAPIResponse struct {
	model.CommonResponse
	AlibabaScbpAdKeywordRankGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaScbpAdKeywordRankGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaScbpAdKeywordRankGetAPIResponseModel).Reset()
}

// AlibabaScbpAdKeywordRankGetAPIResponseModel is 获取外贸直通车关键词预估排名 成功返回结果
type AlibabaScbpAdKeywordRankGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_scbp_ad_keyword_rank_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 关键词的预估排名
	RankLocation int64 `json:"rank_location,omitempty" xml:"rank_location,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaScbpAdKeywordRankGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.RankLocation = 0
}

var poolAlibabaScbpAdKeywordRankGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaScbpAdKeywordRankGetAPIResponse)
	},
}

// GetAlibabaScbpAdKeywordRankGetAPIResponse 从 sync.Pool 获取 AlibabaScbpAdKeywordRankGetAPIResponse
func GetAlibabaScbpAdKeywordRankGetAPIResponse() *AlibabaScbpAdKeywordRankGetAPIResponse {
	return poolAlibabaScbpAdKeywordRankGetAPIResponse.Get().(*AlibabaScbpAdKeywordRankGetAPIResponse)
}

// ReleaseAlibabaScbpAdKeywordRankGetAPIResponse 将 AlibabaScbpAdKeywordRankGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaScbpAdKeywordRankGetAPIResponse(v *AlibabaScbpAdKeywordRankGetAPIResponse) {
	v.Reset()
	poolAlibabaScbpAdKeywordRankGetAPIResponse.Put(v)
}
