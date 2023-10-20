package scbp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpAdKeywordRankPriceBatchgetAPIResponse 外贸直通车关键词前五名批量排价 API返回值
// alibaba.scbp.ad.keyword.rank.price.batchget
//
// 外贸直通车关键词前五名批量排价
type AlibabaScbpAdKeywordRankPriceBatchgetAPIResponse struct {
	model.CommonResponse
	AlibabaScbpAdKeywordRankPriceBatchgetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaScbpAdKeywordRankPriceBatchgetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaScbpAdKeywordRankPriceBatchgetAPIResponseModel).Reset()
}

// AlibabaScbpAdKeywordRankPriceBatchgetAPIResponseModel is 外贸直通车关键词前五名批量排价 成功返回结果
type AlibabaScbpAdKeywordRankPriceBatchgetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_scbp_ad_keyword_rank_price_batchget_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	ResultList []KeywordRankPriceDto `json:"result_list,omitempty" xml:"result_list>keyword_rank_price_dto,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaScbpAdKeywordRankPriceBatchgetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResultList = m.ResultList[:0]
}

var poolAlibabaScbpAdKeywordRankPriceBatchgetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaScbpAdKeywordRankPriceBatchgetAPIResponse)
	},
}

// GetAlibabaScbpAdKeywordRankPriceBatchgetAPIResponse 从 sync.Pool 获取 AlibabaScbpAdKeywordRankPriceBatchgetAPIResponse
func GetAlibabaScbpAdKeywordRankPriceBatchgetAPIResponse() *AlibabaScbpAdKeywordRankPriceBatchgetAPIResponse {
	return poolAlibabaScbpAdKeywordRankPriceBatchgetAPIResponse.Get().(*AlibabaScbpAdKeywordRankPriceBatchgetAPIResponse)
}

// ReleaseAlibabaScbpAdKeywordRankPriceBatchgetAPIResponse 将 AlibabaScbpAdKeywordRankPriceBatchgetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaScbpAdKeywordRankPriceBatchgetAPIResponse(v *AlibabaScbpAdKeywordRankPriceBatchgetAPIResponse) {
	v.Reset()
	poolAlibabaScbpAdKeywordRankPriceBatchgetAPIResponse.Put(v)
}
