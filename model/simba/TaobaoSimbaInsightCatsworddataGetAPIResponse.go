package simba

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSimbaInsightCatsworddataGetAPIResponse
获取类目下关键词的数据 API返回值
taobao.simba.insight.catsworddata.get

获取给定词在给定类目下的详细数据 */
type TaobaoSimbaInsightCatsworddataGetAPIResponse struct {
	model.CommonResponse
	TaobaoSimbaInsightCatsworddataGetAPIResponseModel
}

// TaobaoSimbaInsightCatsworddataGetAPIResponseModel is 获取类目下关键词的数据 成功返回结果
type TaobaoSimbaInsightCatsworddataGetAPIResponseModel struct {
	XMLName xml.Name `xml:"simba_insight_catsworddata_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 关键词在类目下的数据
	CatwordDataList []InsightWordDataUnderCatDto `json:"catword_data_list,omitempty" xml:"catword_data_list>insight_word_data_under_cat_dto,omitempty"`
}
