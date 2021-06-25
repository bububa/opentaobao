package simba

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取类目下关键词的数据 APIResponse
taobao.simba.insight.catsworddata.get

获取给定词在给定类目下的详细数据
*/
type TaobaoSimbaInsightCatsworddataGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoSimbaInsightCatsworddataGetResponse `json:"taobao_simba_insight_catsworddata_get_response,omitempty"`
}

type TaobaoSimbaInsightCatsworddataGetResponse struct {

    // 关键词在类目下的数据
    CatwordDataList   []InsightWordDataUnderCatDTO `json:"catword_data_list,omitempty"`

}
