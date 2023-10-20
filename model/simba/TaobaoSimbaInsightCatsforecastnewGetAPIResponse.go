package simba

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaInsightCatsforecastnewGetAPIResponse 获取词的相关类目预测数据 API返回值
// taobao.simba.insight.catsforecastnew.get
//
// 根据给定的词，预测这些词的相关类目
type TaobaoSimbaInsightCatsforecastnewGetAPIResponse struct {
	model.CommonResponse
	TaobaoSimbaInsightCatsforecastnewGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoSimbaInsightCatsforecastnewGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoSimbaInsightCatsforecastnewGetAPIResponseModel).Reset()
}

// TaobaoSimbaInsightCatsforecastnewGetAPIResponseModel is 获取词的相关类目预测数据 成功返回结果
type TaobaoSimbaInsightCatsforecastnewGetAPIResponseModel struct {
	XMLName xml.Name `xml:"simba_insight_catsforecastnew_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 词的相关类目列表
	CategoryForecastList []InsightCategoryForcastDto `json:"category_forecast_list,omitempty" xml:"category_forecast_list>insight_category_forcast_dto,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoSimbaInsightCatsforecastnewGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.CategoryForecastList = m.CategoryForecastList[:0]
}

var poolTaobaoSimbaInsightCatsforecastnewGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoSimbaInsightCatsforecastnewGetAPIResponse)
	},
}

// GetTaobaoSimbaInsightCatsforecastnewGetAPIResponse 从 sync.Pool 获取 TaobaoSimbaInsightCatsforecastnewGetAPIResponse
func GetTaobaoSimbaInsightCatsforecastnewGetAPIResponse() *TaobaoSimbaInsightCatsforecastnewGetAPIResponse {
	return poolTaobaoSimbaInsightCatsforecastnewGetAPIResponse.Get().(*TaobaoSimbaInsightCatsforecastnewGetAPIResponse)
}

// ReleaseTaobaoSimbaInsightCatsforecastnewGetAPIResponse 将 TaobaoSimbaInsightCatsforecastnewGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoSimbaInsightCatsforecastnewGetAPIResponse(v *TaobaoSimbaInsightCatsforecastnewGetAPIResponse) {
	v.Reset()
	poolTaobaoSimbaInsightCatsforecastnewGetAPIResponse.Put(v)
}
