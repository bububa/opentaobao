package simba

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaKeywordsRecommendGetAPIResponse 取得一个推广组的推荐关键词列表 API返回值
// taobao.simba.keywords.recommend.get
//
// 取得一个推广组的推荐关键词列表
type TaobaoSimbaKeywordsRecommendGetAPIResponse struct {
	model.CommonResponse
	TaobaoSimbaKeywordsRecommendGetAPIResponseModel
}

// TaobaoSimbaKeywordsRecommendGetAPIResponseModel is 取得一个推广组的推荐关键词列表 成功返回结果
type TaobaoSimbaKeywordsRecommendGetAPIResponseModel struct {
	XMLName xml.Name `xml:"simba_keywords_recommend_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 推荐词分页对象，当输入的页码大于最大数值时，将返回最大的page_no值，并且结果中的数据列表为空值
	RecommendWords *RecommendWordPage `json:"recommend_words,omitempty" xml:"recommend_words,omitempty"`
}
