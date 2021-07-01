package simba

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSimbaInsightWordspricedataGetAPIResponse
关键词按竞价区间的分布数据 API返回值
taobao.simba.insight.wordspricedata.get

获取关键词按竞价区间进行细分的数据 */
type TaobaoSimbaInsightWordspricedataGetAPIResponse struct {
	model.CommonResponse
	TaobaoSimbaInsightWordspricedataGetAPIResponseModel
}

// TaobaoSimbaInsightWordspricedataGetAPIResponseModel is 关键词按竞价区间的分布数据 成功返回结果
type TaobaoSimbaInsightWordspricedataGetAPIResponseModel struct {
	XMLName xml.Name `xml:"simba_insight_wordspricedata_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 竞价区间分布数据
	WordPricedataList []InsightWordPriceDistributeDataDto `json:"word_pricedata_list,omitempty" xml:"word_pricedata_list>insight_word_price_distribute_data_dto,omitempty"`
}
