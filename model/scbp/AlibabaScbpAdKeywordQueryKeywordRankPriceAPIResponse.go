package scbp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpAdKeywordQueryKeywordRankPriceAPIResponse 查询关键词前五名排价 API返回值
// alibaba.scbp.ad.keyword.query.keyword.rank.price
//
// 查询关键词前五名排价
type AlibabaScbpAdKeywordQueryKeywordRankPriceAPIResponse struct {
	model.CommonResponse
	AlibabaScbpAdKeywordQueryKeywordRankPriceAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaScbpAdKeywordQueryKeywordRankPriceAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaScbpAdKeywordQueryKeywordRankPriceAPIResponseModel).Reset()
}

// AlibabaScbpAdKeywordQueryKeywordRankPriceAPIResponseModel is 查询关键词前五名排价 成功返回结果
type AlibabaScbpAdKeywordQueryKeywordRankPriceAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_scbp_ad_keyword_query_keyword_rank_price_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回参数
	Result *KeywordRankPriceDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaScbpAdKeywordQueryKeywordRankPriceAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaScbpAdKeywordQueryKeywordRankPriceAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaScbpAdKeywordQueryKeywordRankPriceAPIResponse)
	},
}

// GetAlibabaScbpAdKeywordQueryKeywordRankPriceAPIResponse 从 sync.Pool 获取 AlibabaScbpAdKeywordQueryKeywordRankPriceAPIResponse
func GetAlibabaScbpAdKeywordQueryKeywordRankPriceAPIResponse() *AlibabaScbpAdKeywordQueryKeywordRankPriceAPIResponse {
	return poolAlibabaScbpAdKeywordQueryKeywordRankPriceAPIResponse.Get().(*AlibabaScbpAdKeywordQueryKeywordRankPriceAPIResponse)
}

// ReleaseAlibabaScbpAdKeywordQueryKeywordRankPriceAPIResponse 将 AlibabaScbpAdKeywordQueryKeywordRankPriceAPIResponse 保存到 sync.Pool
func ReleaseAlibabaScbpAdKeywordQueryKeywordRankPriceAPIResponse(v *AlibabaScbpAdKeywordQueryKeywordRankPriceAPIResponse) {
	v.Reset()
	poolAlibabaScbpAdKeywordQueryKeywordRankPriceAPIResponse.Put(v)
}
