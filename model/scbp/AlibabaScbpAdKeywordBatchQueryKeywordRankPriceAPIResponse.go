package scbp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpAdKeywordBatchQueryKeywordRankPriceAPIResponse 批量查询关键词前五名排价 API返回值
// alibaba.scbp.ad.keyword.batch.query.keyword.rank.price
//
// 批量查询关键词前五名排价
type AlibabaScbpAdKeywordBatchQueryKeywordRankPriceAPIResponse struct {
	model.CommonResponse
	AlibabaScbpAdKeywordBatchQueryKeywordRankPriceAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaScbpAdKeywordBatchQueryKeywordRankPriceAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaScbpAdKeywordBatchQueryKeywordRankPriceAPIResponseModel).Reset()
}

// AlibabaScbpAdKeywordBatchQueryKeywordRankPriceAPIResponseModel is 批量查询关键词前五名排价 成功返回结果
type AlibabaScbpAdKeywordBatchQueryKeywordRankPriceAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_scbp_ad_keyword_batch_query_keyword_rank_price_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 关键词前五名排价详细信息返回
	ResultList []KeywordRankPriceDto `json:"result_list,omitempty" xml:"result_list>keyword_rank_price_dto,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaScbpAdKeywordBatchQueryKeywordRankPriceAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResultList = m.ResultList[:0]
}

var poolAlibabaScbpAdKeywordBatchQueryKeywordRankPriceAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaScbpAdKeywordBatchQueryKeywordRankPriceAPIResponse)
	},
}

// GetAlibabaScbpAdKeywordBatchQueryKeywordRankPriceAPIResponse 从 sync.Pool 获取 AlibabaScbpAdKeywordBatchQueryKeywordRankPriceAPIResponse
func GetAlibabaScbpAdKeywordBatchQueryKeywordRankPriceAPIResponse() *AlibabaScbpAdKeywordBatchQueryKeywordRankPriceAPIResponse {
	return poolAlibabaScbpAdKeywordBatchQueryKeywordRankPriceAPIResponse.Get().(*AlibabaScbpAdKeywordBatchQueryKeywordRankPriceAPIResponse)
}

// ReleaseAlibabaScbpAdKeywordBatchQueryKeywordRankPriceAPIResponse 将 AlibabaScbpAdKeywordBatchQueryKeywordRankPriceAPIResponse 保存到 sync.Pool
func ReleaseAlibabaScbpAdKeywordBatchQueryKeywordRankPriceAPIResponse(v *AlibabaScbpAdKeywordBatchQueryKeywordRankPriceAPIResponse) {
	v.Reset()
	poolAlibabaScbpAdKeywordBatchQueryKeywordRankPriceAPIResponse.Put(v)
}
