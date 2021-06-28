package simba

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取词的相关类目预测数据 APIResponse
taobao.simba.insight.catsforecastnew.get

根据给定的词，预测这些词的相关类目
*/
type TaobaoSimbaInsightCatsforecastnewGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoSimbaInsightCatsforecastnewGetResponse `json:"simba_insight_catsforecastnew_get_response,omitempty"` 
    TaobaoSimbaInsightCatsforecastnewGetResponse
}

/* model for simplify = false
type TaobaoSimbaInsightCatsforecastnewGetResponse struct {

    // 词的相关类目列表
    
    CategoryForecastList  struct {
        InsightCategoryForcastDTO  []InsightCategoryForcastDTO `json:"insight_category_forcast_dto,omitempty"`
    } `json:"category_forecast_list,omitempty"`
    

}
*/

type TaobaoSimbaInsightCatsforecastnewGetResponse struct {

    // 词的相关类目列表
    CategoryForecastList   []InsightCategoryForcastDTO `json:"category_forecast_list,omitempty"`

}
