package simba

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取词的相关类目预测数据 API返回值 
taobao.simba.insight.catsforecastnew.get

根据给定的词，预测这些词的相关类目
*/
type TaobaoSimbaInsightCatsforecastnewGetAPIResponse struct {
    model.CommonResponse
    TaobaoSimbaInsightCatsforecastnewGetResponse
}

// 获取词的相关类目预测数据 成功返回结果
type TaobaoSimbaInsightCatsforecastnewGetResponse struct {
    XMLName xml.Name `xml:"simba_insight_catsforecastnew_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 词的相关类目列表
    CategoryForecastList   []InsightCategoryForcastDTO `json:"category_forecast_list,omitempty" xml:"category_forecast_list>insight_category_forcast_dto,omitempty"`
}
