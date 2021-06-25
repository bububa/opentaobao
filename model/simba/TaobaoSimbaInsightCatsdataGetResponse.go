package simba

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取类目的大盘数据 APIResponse
taobao.simba.insight.catsdata.get

根据类目id获取类目的大盘数据，包括展现指数，点击指数，点击率，本次提供的insight相关的其它接口的都是这种情况。
*/
type TaobaoSimbaInsightCatsdataGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoSimbaInsightCatsdataGetResponse `json:"taobao_simba_insight_catsdata_get_response,omitempty"`
}

type TaobaoSimbaInsightCatsdataGetResponse struct {

    // 类目详细数据列表
    CatDataList   []InsightCategoryDataDTO `json:"cat_data_list,omitempty"`

}
