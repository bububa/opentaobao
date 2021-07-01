package scbp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaScbpAdKeywordRankPriceGetAPIResponse
外贸直通车关键词前五名排价 API返回值
alibaba.scbp.ad.keyword.rank.price.get

外贸直通车关键词前五名排价 */
type AlibabaScbpAdKeywordRankPriceGetAPIResponse struct {
	model.CommonResponse
	AlibabaScbpAdKeywordRankPriceGetAPIResponseModel
}

// AlibabaScbpAdKeywordRankPriceGetAPIResponseModel is 外贸直通车关键词前五名排价 成功返回结果
type AlibabaScbpAdKeywordRankPriceGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_scbp_ad_keyword_rank_price_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 关键词前五名排价
	RankPriceList []string `json:"rank_price_list,omitempty" xml:"rank_price_list>string,omitempty"`
}
