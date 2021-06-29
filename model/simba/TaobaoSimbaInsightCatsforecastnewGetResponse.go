package simba

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取词的相关类目预测数据 APIResponse
taobao.simba.insight.catsforecastnew.get

根据给定的词，预测这些词的相关类目
*/
type TaobaoSimbaInsightCatsforecastnewGetAPIResponse struct {
    model.CommonResponse
    TaobaoSimbaInsightCatsforecastnewGetResponse
}

type TaobaoSimbaInsightCatsforecastnewGetResponse struct {
    XMLName xml.Name `xml:"simba_insight_catsforecastnew_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 词的相关类目列表
    
    CategoryForecastList   []InsightCategoryForcastDTO `json:"category_forecast_list,omitempty" xml:"category_forecast_list>insight_category_forcast_dto,omitempty"`
    
    
}
