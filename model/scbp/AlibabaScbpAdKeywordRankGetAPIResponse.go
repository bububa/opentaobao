package scbp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaScbpAdKeywordRankGetAPIResponse
获取外贸直通车关键词预估排名 API返回值
alibaba.scbp.ad.keyword.rank.get

获取外贸直通车关键词预估排名 */
type AlibabaScbpAdKeywordRankGetAPIResponse struct {
	model.CommonResponse
	AlibabaScbpAdKeywordRankGetAPIResponseModel
}

// AlibabaScbpAdKeywordRankGetAPIResponseModel is 获取外贸直通车关键词预估排名 成功返回结果
type AlibabaScbpAdKeywordRankGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_scbp_ad_keyword_rank_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 关键词的预估排名
	RankLocation int64 `json:"rank_location,omitempty" xml:"rank_location,omitempty"`
}
