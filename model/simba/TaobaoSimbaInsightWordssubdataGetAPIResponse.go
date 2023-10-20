package simba

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaInsightWordssubdataGetAPIResponse 获取关键词按流量细分的数据 API返回值
// taobao.simba.insight.wordssubdata.get
//
// 获取关键词按流量进行细分的数据，返回结果中network表示流量的来源，意义如下：1-&gt;PC站内，2-&gt;PC站外,4-&gt;无线站内 5-&gt;无线站外
type TaobaoSimbaInsightWordssubdataGetAPIResponse struct {
	model.CommonResponse
	TaobaoSimbaInsightWordssubdataGetAPIResponseModel
}

// TaobaoSimbaInsightWordssubdataGetAPIResponseModel is 获取关键词按流量细分的数据 成功返回结果
type TaobaoSimbaInsightWordssubdataGetAPIResponseModel struct {
	XMLName xml.Name `xml:"simba_insight_wordssubdata_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 关键词按流量细分的数据
	WordSubdataList []InsightWordSubDataDto `json:"word_subdata_list,omitempty" xml:"word_subdata_list>insight_word_sub_data_dto,omitempty"`
}
